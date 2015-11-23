package dinheiro

import (
	"errors"
	"math"
)

const (
	negativeValueError = "Não é possível transformar números negativos."
	unsupportedValueError      = "Número muito grande para ser transformado em extenso."

	andSeparator = " e "

	currencyCentavo = "centavo"
	currencyCentavos  = "centavos"
	currencyReal = "real"
	currencyReais   = "reais"
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
	cem      = "cem"
	milhar   = "mil"
	milhao   = "milhão"
	milhoes  = "milhôes"
	bilhao   = "bilhão"
	bilhoes  = "bilhões"
)

// Real é a moeda corrente no Brasil
// en: Real is the present-day currency of Brazil
type Real float64

// PorExtenso Retorna o value por extenso do dinheiro
// en: Returns the value into words
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

func (real Real) String() (string, error) {
	return real.PorExtenso()
}

func convertNumberIntoWords(f float64) (string, error) {
	switch {
	case f < 0:
		return "", errors.New(negativeValueError)
	case f < 20:
		return numbers[int(f)], nil
	case f >= 20 && f < 100:
		value := tens[int((f-20)/10)]
		mod := math.Mod(f, 10)
		if mod != 0 {
			complemento, _ := convertNumberIntoWords(mod)
			value += " e " + complemento
		}
		return value, nil
	case f >= 100 && f < 101:
		return cem, nil
	case f >= 101 && f < 1000:
		value := hundreds[int(f/100-1)]
		mod := math.Mod(f, 100)
		if mod != 0 {
			complemento, _ := convertNumberIntoWords(mod)
			value += " e " + complemento
		}
		return value, nil
	case f == 1000:
		return milhar, nil
	default:
		return "", errors.New(unsupportedValueError)
	}
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
