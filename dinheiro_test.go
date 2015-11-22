package dinheiro

import "testing"

func TestValorNegativo(t *testing.T) {
	var real Real
	real = -1
	_, err := real.PorExtenso()

	if err == nil {
		t.Errorf("Expected '%v' got '%v'", valorNegativoException, err.Error())
	}
}

func TestValorSuperiorLimite(t *testing.T) {
	var real Real
	real = 99999999999999
	_, err := real.PorExtenso()

	if err == nil {
		t.Errorf("Expected '%v' got '%v'", valorNaoSuportado, err.Error())
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
	var valoresPorExtenso = [36]string{
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
		want := valoresPorExtenso[i]
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

func TestValoresInteiros(t *testing.T) {
	var valoresPorExtenso = [20]string{
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
		want := valoresPorExtenso[i]
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

func TestDezenas(t *testing.T) {
	dezenas := [8]string{
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
		want := dezenas[i]

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

func TestCentenas(t *testing.T) {
	dezenas := [9]string{
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
		want := dezenas[i-1]

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
