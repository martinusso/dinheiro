package dinheiro

import (
	"errors"
	"math"
	"strings"
)

const (
	negativeValueError = "Não é possível transformar números negativos."

	andSeparator = "e"
	comma        = ","

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

	thousandsSingular = []string{"", "mil", "milhão", "bilhão", "trilhão", "quadrilhão",
		"quintilhão", "sextilhão", "septilhão", "octilhão", "nonilhão", "decilhão",
		"undecilhão", "duodecilhão", "tredecilhão", "quattuordecilhão", "quindecilhão",
		"sexdecilhão", "septendecilhão", "octodecilhão", "novendecilhão", "vigintilhão"}

	thousands = []string{"", "mil", "milhões", "bilhões", "trilhão", "quadrilhões",
		"quintilhões", "sextilhões", "septilhões", "octilhões", "nonilhões", "decilhões",
		"undecilhões", "duodecilhões", "tredecilhões", "quattuordecilhões", "quindecilhões",
		"sexdecilhões", "septendecilhões", "octodecilhões", "novendecilhões", "vigintilhões"}
)

// Real é a moeda corrente no Brasil
// en: Real is the present-day currency of Brazil
type Real float64

// PorExtenso Retorna o value por extenso do dinheiro
// en: Returns the value into words
func (real Real) PorExtenso() (string, error) {
	words := []string{}

	integer, fractional := math.Modf(float64(real))

	if integer != 0 || fractional == 0 {
		numberIntoWords, err := writeOutNumbersInWords(integer)
		if err != nil {
			return "", err
		}
		words = append(words, numberIntoWords...)
		words = append(words, getIntegerUnit(integer))

	}

	if fractional > 0 {
		fractional := round(math.Abs(fractional) * 100)
		numberIntoWords, err := writeOutNumbersInWords(fractional)
		if err != nil {
			return "", err
		}
		if integer > 0 {
			words = append(words, andSeparator)
		}
		words = append(words, numberIntoWords...)
		words = append(words, getDecimalUnit(fractional))
	}
	return sanitize(words), nil
}

func writeOutNumbersInWords(f float64) ([]string, error) {
	switch {
	case f < 0:
		return []string{}, errors.New(negativeValueError)
	case f < 20:
		s := numbers[int(f)]
		return []string{s}, nil
	case f < 100:
		return getNumberUnderHundred(f)
	case f == 100:
		return []string{hundred}, nil
	case f < 1000:
		return getNumberUnderThousand(f)
	case f == 1000:
		return []string{thousand}, nil
	default:
		return getUpThousand(f)
	}
}

func getNumberUnderHundred(f float64) ([]string, error) {
	value := tens[int((f-20)/10)]
	words := []string{value}

	mod := math.Mod(f, 10)
	if mod != 0 {
		words = append(words, andSeparator)
		words = append(words, numbers[int(mod)])
	}
	return words, nil
}

func getNumberUnderThousand(f float64) ([]string, error) {
	value := hundreds[int(f/100-1)]
	words := []string{value}

	mod := math.Mod(f, 100)
	if mod != 0 {
		remaining, _ := writeOutNumbersInWords(mod)
		words = append(words, andSeparator)
		words = append(words, remaining...)
	}
	return words, nil
}

func getUpThousand(f float64) (words []string, err error) {
	for i := 0; f >= 1; i++ {
		var r float64
		f, r = math.Modf(f / 1000)
		r = round(math.Abs(r) * 1000)

		w, _ := writeOutNumbersInWords(r)
		if (len(w) == 0) || (len(w) == 1 && w[0] == "") {
			continue
		}

		if i > 0 {
			w = append(w, getThousands(r, i))

			if i > 1 {
				w = append(w, comma)
			} else {
				w = append(w, andSeparator)
			}
		}
		words = append(w, words...)
	}
	return words, nil
}

func getThousands(value float64, index int) string {
	if value > 1 {
		return thousands[index]
	}
	return thousandsSingular[index]
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

func sanitize(words []string) string {
	s := strings.Join(words, " ")
	s = strings.Replace(s, " , ", ", ", -1)
	s = strings.Trim(s, " ")
	return s
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
