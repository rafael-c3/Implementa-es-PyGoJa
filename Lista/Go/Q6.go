package main

import "fmt"

type Motor struct {
    Potencia int
    Tipo     string
}

func (m Motor) ExibirDetalhes() {
    fmt.Printf("Motor: %s com %d cavalos de potência\n", m.Tipo, m.Potencia)
}

type Carro struct {
    Marca  string
    Modelo string
    Ano    int
    Motor  Motor
}

func (c Carro) ExibirDetalhes() {
    fmt.Printf("Carro: %s %s, Ano: %d\n", c.Marca, c.Modelo, c.Ano)
    c.Motor.ExibirDetalhes()
}

func main() {
    motor1 := Motor{
        Potencia: 180,
        Tipo:     "V8",
    }

    carro1 := Carro{
        Marca:  "Ford",
        Modelo: "Mustang",
        Ano:    2022,
        Motor:  motor1,
    }

    carro1.ExibirDetalhes()
}
