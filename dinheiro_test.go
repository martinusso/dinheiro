package dinheiro

import "testing"

func TestRealPorExtenso(t *testing.T) {
	values := map[float64]string{
		1.99:       "um real e noventa e nove centavos",
		42.69:      "quarenta e dois reais e sessenta e nove centavos",
		2015:       "dois mil e quinze reais",
		802715:     "oitocentos e dois mil e setecentos e quinze reais",
		775398007:  "setecentos e setenta e cinco milhões, trezentos e noventa e oito mil e sete reais",
		1119929521: "um bilhão, cento e dezenove milhões, novecentos e vinte e nove mil e quinhentos e vinte e um reais",
	}
	for f, expected := range values {
		got, err := RealPorExtenso(f)

		if err != nil {
			t.Errorf("Error should be nil => %s", err)
		}
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}

func TestNegativeValue(t *testing.T) {
	var real Real
	real = -1
	_, err := real.PorExtenso()

	if err == nil {
		t.Errorf("Expected '%v' got '%v'", negativeValueError, nil)
	}
}

func TestCentavos(t *testing.T) {
	valores := map[float64]string{
		0.01234: "um centavo",
		0.01:    "um centavo",
		0.02:    "dois centavos",
		0.02345: "dois centavos",
		0.03:    "três centavos",
		0.04:    "quatro centavos",
		0.05:    "cinco centavos",
		0.06:    "seis centavos",
		0.07:    "sete centavos",
		0.08:    "oito centavos",
		0.09:    "nove centavos",
		0.1:     "dez centavos",
		0.11:    "onze centavos",
		0.12:    "doze centavos",
		0.13:    "treze centavos",
		0.14:    "quatorze centavos",
		0.15:    "quinze centavos",
		0.16:    "dezesseis centavos",
		0.17:    "dezessete centavos",
		0.18:    "dezoito centavos",
		0.19:    "dezenove centavos",
		0.2:     "vinte centavos",
		0.22:    "vinte e dois centavos",
		0.3:     "trinta centavos",
		0.33:    "trinta e três centavos",
		0.4:     "quarenta centavos",
		0.44:    "quarenta e quatro centavos",
		0.5:     "cinquenta centavos",
		0.55:    "cinquenta e cinco centavos",
		0.6:     "sessenta centavos",
		0.66:    "sessenta e seis centavos",
		0.7:     "setenta centavos",
		0.77:    "setenta e sete centavos",
		0.8:     "oitenta centavos",
		0.88:    "oitenta e oito centavos",
		0.9:     "noventa centavos",
		0.99:    "noventa e nove centavos",
	}
	var real Real
	for f, expected := range valores {
		real = Real(f)
		got, err := real.PorExtenso()

		if err != nil {
			t.Errorf("Expected '%v' got '%v'", nil, err)
		}
		if got != expected {
			t.Errorf("Expected '%v' got '%v' - '%f'", expected, got, real)
		}
	}
}

func TestIntegersValues(t *testing.T) {
	var valuesIntoWords = [20]string{
		"zero real",
		"um real",
		"dois reais",
		"três reais",
		"quatro reais",
		"cinco reais",
		"seis reais",
		"sete reais",
		"oito reais",
		"nove reais",
		"dez reais",
		"onze reais",
		"doze reais",
		"treze reais",
		"quatorze reais",
		"quinze reais",
		"dezesseis reais",
		"dezessete reais",
		"dezoito reais",
		"dezenove reais",
	}
	var real Real
	for i := 0; i < 19; i++ {
		want := valuesIntoWords[i]
		real = Real(i)
		got, err := real.PorExtenso()

		if err != nil {
			t.Errorf("Expected '%v' got '%v'", nil, err)
		}

		if got != want {
			t.Errorf("Expected '%v' got '%v'", want, got)
		}
	}
}

func TestTens(t *testing.T) {
	tens := [8]string{
		"vinte reais",
		"trinta reais",
		"quarenta reais",
		"cinquenta reais",
		"sessenta reais",
		"setenta reais",
		"oitenta reais",
		"noventa reais",
	}
	var real Real
	for i := 0; i < 8; i++ {
		want := tens[i]

		real = Real(i*10 + 20)
		got, err := real.PorExtenso()

		if err != nil {
			t.Errorf("Expected '%v' got '%v'", nil, err)
		}

		if got != want {
			t.Errorf("Expected '%v' got '%v'", want, got)
		}
	}
}

func TestHundreds(t *testing.T) {
	hundreds := [9]string{
		"cem reais",
		"duzentos reais",
		"trezentos reais",
		"quatrocentos reais",
		"quinhentos reais",
		"seiscentos reais",
		"setecentos reais",
		"oitocentos reais",
		"novecentos reais",
	}
	var real Real
	for i := 1; i <= 9; i++ {
		want := hundreds[i-1]

		real = Real(i * 100)
		got, err := real.PorExtenso()

		if err != nil {
			t.Errorf("Expected '%v' got '%v'", nil, err)
		}

		if got != want {
			t.Errorf("Expected '%v' got '%v'", want, got)
		}
	}
}

func TestMilReais(t *testing.T) {
	want := "mil reais"
	var real Real
	real = 1000
	got, err := real.PorExtenso()

	if err != nil {
		t.Errorf("Expected '%v' got '%v'", nil, err)
	}

	if got != want {
		t.Errorf("Expected '%v' got '%v'", want, got)
	}
}
