package dinheiro

import "testing"

func TestNegativeValue(t *testing.T) {
	var real Real
	real = -1
	_, err := real.PorExtenso()

	if err == nil {
		t.Errorf("Expected '%v' got '%v'", negativeValueError, nil)
	}
}

func TestCentavos(t *testing.T) {
	var valores = [36]float64{
		0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07, 0.08, 0.09,
		0.1, 0.11, 0.12, 0.13, 0.14, 0.15, 0.16, 0.17, 0.18, 0.19,
		0.20, 0.22,
		0.30, 0.33,
		0.40, 0.44,
		0.50, 0.55,
		0.60, 0.66,
		0.70, 0.77,
		0.80, 0.88,
		0.90, 0.99,
	}
	var valuesIntoWords = [36]string{
		"um centavo",
		"dois centavos",
		"três centavos",
		"quatro centavos",
		"cinco centavos",
		"seis centavos",
		"sete centavos",
		"oito centavos",
		"nove centavos",
		"dez centavos",
		"onze centavos",
		"doze centavos",
		"treze centavos",
		"quatorze centavos",
		"quinze centavos",
		"dezesseis centavos",
		"dezessete centavos",
		"dezoito centavos",
		"dezenove centavos",
		"vinte centavos",
		"vinte e dois centavos",
		"trinta centavos",
		"trinta e três centavos",
		"quarenta centavos",
		"quarenta e quatro centavos",
		"cinquenta centavos",
		"cinquenta e cinco centavos",
		"sessenta centavos",
		"sessenta e seis centavos",
		"setenta centavos",
		"setenta e sete centavos",
		"oitenta centavos",
		"oitenta e oito centavos",
		"noventa centavos",
		"noventa e nove centavos",
	}
	var real Real
	for i := 0; i < 35; i++ {
		want := valuesIntoWords[i]
		real = Real(valores[i])
		got, err := real.PorExtenso()

		if err != nil {
			t.Errorf("Expected '%v' got '%v'", nil, err)
		}

		if got != want {
			t.Errorf("Expected '%v' got '%v' - '%f'", want, got, real)
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

func TestUmRealENoventaENoveCentavos(t *testing.T) {
	want := "um real e noventa e nove centavos"
	var real Real
	real = 1.99
	got, err := real.PorExtenso()

	if err != nil {
		t.Errorf("Expected '%v' got '%v'", nil, err)
	}

	if got != want {
		t.Errorf("Expected '%v' got '%v'", want, got)
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

func TestQuarentaEDoisReaisESessentaENoveCentavos(t *testing.T) {
	want := "quarenta e dois reais e sessenta e nove centavos"
	var real Real
	real = 42.69
	got, err := real.PorExtenso()

	if err != nil {
		t.Errorf("Expected '%v' got '%v'", nil, err)
	}

	if got != want {
		t.Errorf("Expected '%v' got '%v'", want, got)
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

func TestDoisMilEQuinzeReais(t *testing.T) {
	want := "dois mil e quinze reais"
	var real Real
	real = 2015
	got, err := real.PorExtenso()

	if err != nil {
		t.Errorf("Expected '%v' got '%v'", nil, err)
	}

	if got != want {
		t.Errorf("Expected '%v' got '%v'", want, got)
	}
}

func TestOitocentosEDoisMilESetencentosQuinzeReais(t *testing.T) {
	want := "oitocentos e dois mil e setecentos e quinze reais"
	var real Real
	real = 802715
	got, err := real.PorExtenso()

	if err != nil {
		t.Errorf("Expected '%v' got '%v'", nil, err)
	}

	if got != want {
		t.Errorf("Expected '%v' got '%v'", want, got)
	}
}

func TestMilhoes(t *testing.T) {
	want := "setecentos e setenta e cinco milhões, trezentos e noventa e oito mil e sete reais"
	var real Real
	real = 775398007
	got, err := real.PorExtenso()

	if err != nil {
		t.Errorf("Expected '%v' got '%v'", nil, err)
	}

	if got != want {
		t.Errorf("Expected '%v' got '%v'", want, got)
	}
}

func TestBilhao(t *testing.T) {
	want := "um bilhão, cento e dezenove milhões, novecentos e vinte e nove mil e quinhentos e vinte e um reais"
	var real Real
	real = 1119929521
	got, err := real.PorExtenso()

	if err != nil {
		t.Errorf("Expected '%v' got '%v'", nil, err)
	}

	if got != want {
		t.Errorf("Expected '%v' got '%v'", want, got)
	}
}
