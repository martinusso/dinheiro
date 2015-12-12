// Package dinheiro Converte um valor monetário (em Real) por extenso
// en: Returns the Real Currency (the present-day currency of Brazil) into words
package dinheiro

import (
	"errors"
	"math"
	"strconv"
)

const (
	negativeValueError    = "Não é possível transformar números negativos."
	unsupportedValueError = "Número muito grande para ser transformado em extenso."

	andSeparator = " e "

	currencyCentavo  = "centavo"
	currencyCentavos = "centavos"
	currencyReal     = "real"
	currencyReais    = "reais"
)

var (
	numbers = [20]string{
		"zero",
		"um",
		"dois",
		"três",
		"quatro",
		"cinco",
		"seis",
		"sete",
		"oito",
		"nove",
		"dez",
		"onze",
		"doze",
		"treze",
		"quatorze",
		"quinze",
		"dezesseis",
		"dezessete",
		"dezoito",
		"dezenove"}

	tens = [8]string{
		"vinte",
		"trinta",
		"quarenta",
		"cinquenta",
		"sessenta",
		"setenta",
		"oitenta",
		"noventa",
	}

	hundreds = [9]string{
		"cento",
		"duzentos",
		"trezentos",
		"quatrocentos",
		"quinhentos",
		"seiscentos",
		"setecentos",
		"oitocentos",
		"novecentos",
	}
	hundred  = "cem"
	thousand = "mil"
	million  = "milhão"
	millions = "milhões"
	billion  = "bilhão"
	billions = "bilhões"
)

// RealPorExtenso Retorna um valor em Real por extenso
// [en: Returns a value (Real Currency as float64) into words]
func RealPorExtenso(value float64) (string, error) {
	return Real(value).PorExtenso()
}

// Real é a moeda corrente no Brasil
// [en: Real is the present-day currency of Brazil]
type Real float64

// PorExtenso Retorna o valor por extenso
// [en: Returns the value into words]
func (real Real) PorExtenso() (string, error) {
	var value string

	integer, fractional := math.Modf(float64(real))

	if integer != 0 || fractional == 0 {
		numberIntoWords, err := convertNumberIntoWords(integer)
		if err != nil {
			return "", err
		}
		value = numberIntoWords + " " + getIntegerUnit(integer)
	}

	if fractional > 0 {
		fractional := round(math.Abs(fractional) * 100)
		numberIntoWords, err := convertNumberIntoWords(fractional)
		if err != nil {
			return "", err
		}
		if integer > 0 {
			value += andSeparator
		}
		value += numberIntoWords + " " + getDecimalUnit(fractional)
	}

	return value, nil
}

func convertNumberIntoWords(f float64) (string, error) {
	switch {
	case f < 0:
		return "", errors.New(negativeValueError)
	case f < 20:
		return numbers[int(f)], nil
	case f < 100:
		return getNumberUnderHundred(f)
	case f == 100:
		return hundred, nil
	case f < 1000:
		return getNumberUnderThousand(f)
	case f == 1000:
		return thousand, nil
	case f < 1000000:
		return getNumberUnderMillion(f)
	default:
		return "", errors.New(unsupportedValueError)
	}
}

func getNumberUnderHundred(f float64) (string, error) {
	value := tens[int((f-20)/10)]
	mod := math.Mod(f, 10)
	if mod != 0 {
		value += andSeparator + numbers[int(mod)]
	}
	return value, nil
}

func getNumberUnderThousand(f float64) (string, error) {
	value := hundreds[int(f/100-1)]
	mod := math.Mod(f, 100)
	if mod != 0 {
		remaining, _ := convertNumberIntoWords(mod)
		value += andSeparator + remaining
	}
	return value, nil
}

func getNumberUnderMillion(f float64) (string, error) {
	s := strconv.Itoa(int(f))
	t1, _ := strconv.Atoi(s[:len(s)-3])
	t2, _ := strconv.Atoi(s[len(s)-3:])

	value, _ := convertNumberIntoWords(float64(t1))
	value += " " + thousand
	if t2 > 0 {
		t2IntoWords, _ := convertNumberIntoWords(float64(t2))
		value += andSeparator + t2IntoWords
	}
	return value, nil
}

func getIntegerUnit(f float64) string {
	if f >= 0 && f < 2 {
		return currencyReal
	}
	return currencyReais
}

func getDecimalUnit(f float64) string {
	if f == 1 {
		return currencyCentavo
	}
	return currencyCentavos
}

func round(val float64) float64 {
	const (
		roundOn = .5
		places  = 2
	)

	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / pow
}
