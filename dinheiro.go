package dinheiro

import (
	"errors"
	"math"
)

const (
	valorInferiorZeroException = "Valor inferior a zero reais"

	e = "e"

	centavoSingular = "centavo"
	centavosPlural  = "centavos"

	realSingular = "real"
	realPlural   = "reais"
)

var (
	numeros = [20]string{
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

	dezenas = [8]string{
		"vinte",
		"trinta",
		"quarenta",
		"cinquenta",
		"sessenta",
		"setenta",
		"oitenta",
		"noventa",
	}

	centenas = [9]string{
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
	trilhao  = "trilhão"
	trilhoes = "trilhôes"
)

// Dinheiro interface do dinheiro
type Dinheiro interface {
	PorExtenso() (string, error)
}

// Real é a moeda corrente no Brasil
type Real float64

// PorExtenso Retorna o valor por extenso do dinheiro
func (real Real) PorExtenso() (string, error) {
	inteiro, fracionario := math.Modf(float64(real))

	var valor string

	if inteiro != 0 || fracionario == 0 {
		numeroPorExtenso, err := getNumero(inteiro)
		if err != nil {
			return "", err
		}
		moeda := getMoeda(float64(inteiro))
		valor = numeroPorExtenso + " " + moeda
	}

	if fracionario > 0 {
		fracionario := round(math.Abs(fracionario) * 100)
		numeroPorExtenso, err := getNumero(fracionario)
		if err != nil {
			return "", err
		}
		if inteiro > 0 {
			valor += " e "
		}
		valor += numeroPorExtenso + " " + getCents(fracionario)
	}

	return valor, nil
}

func (real Real) String() (string, error) {
	return real.PorExtenso()
}

func getNumero(f float64) (string, error) {
	switch {
	case f < 0:
		return "", errors.New(valorInferiorZeroException)
	case f < 20:
		return numeros[int(f)], nil
	case f >= 20 && f < 100:
		valor := dezenas[int((f-20)/10)]
		mod := math.Mod(f, 10)
		if mod != 0 {
			complemento, _ := getNumero(mod)
			valor += " e " + complemento
		}
		return valor, nil
	case f >= 100 && f < 101:
		return cem, nil
	case f >= 101 && f < 1000:
		valor := centenas[int(f/100-1)]
		mod := math.Mod(f, 100)
		if mod != 0 {
			complemento, _ := getNumero(mod)
			valor += " e " + complemento
		}
		return valor, nil
	case f == 1000:
		return "mil", nil
	default:
		return "", errors.New("Valor não suportado")
	}
}

func getMoeda(f float64) string {
	if f >= 0 && f < 2 {
		return realSingular
	}
	return realPlural
}

func getCents(f float64) string {
	if f == 1 {
		return centavoSingular
	}
	return centavosPlural
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
