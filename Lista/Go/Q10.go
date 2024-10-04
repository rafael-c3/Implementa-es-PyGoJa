package main

import "fmt"

type Calculadora struct{}

func (c Calculadora) SomarDois(a, b float64) float64 {
    return a + b
}

func (c Calculadora) SomarTres(a, b, c1 float64) float64 {
    return a + b + c1
}

func main() {
    calc := Calculadora{}

    resultadoDois := calc.SomarDois(10, 20)
    fmt.Printf("Soma de dois números: %.2f\n", resultadoDois)

    resultadoTres := calc.SomarTres(10, 20, 30)
    fmt.Printf("Soma de três números: %.2f\n", resultadoTres)
}
