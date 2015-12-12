# dinheiro

Converte um valor monetário (em Real) por extenso


*[en]: Returns the [Real](https://en.wikipedia.org/wiki/Brazilian_real) (the present-day currency of Brazil) into words. The return will be in Portuguese.*

[![Circle CI](https://circleci.com/gh/martinusso/dinheiro/tree/master.svg?style=shield&circle-token=:circle-token)](https://circleci.com/gh/martinusso/dinheiro/tree/master)
[![GoDoc](https://godoc.org/github.com/martinusso/dinheiro?status.svg)](https://godoc.org/github.com/martinusso/dinheiro)

## Usage

### RealPorExtenso
```go
value := 1.99
valueIntoWords, _ := RealPorExtenso(value)
// valueIntoWords = "um real e noventa e nove centavos"
```

### Using the Real type
```go
real := Real(1.99)
realIntoWords, _ := real.PorExtenso()
// realIntoWords = "um real e noventa e nove centavos"
```

## License

This software is open source, licensed under the The MIT License (MIT). See [LICENSE](https://github.com/martinusso/dinheiro/blob/master/LICENSE) for details.
