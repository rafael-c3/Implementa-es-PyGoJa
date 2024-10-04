package main

import "fmt"

type Carro struct {
    Marca      string
    Modelo     string
    Ano        int
    Velocidade int
}

func (c Carro) ExibirDetalhes() {
    fmt.Printf("Marca: %s, Modelo: %s, Ano: %d\n", c.Marca, c.Modelo, c.Ano)
}

func (c *Carro) Acelerar(valor int) {
    c.Velocidade += valor
    fmt.Printf("%s %s acelerou para %d km/h\n", c.Marca, c.Modelo, c.Velocidade)
}

func (c *Carro) Frear(valor int) {
    c.Velocidade -= valor
    if c.Velocidade < 0 {
        c.Velocidade = 0
    }
    fmt.Printf("%s %s reduziu a velocidade para %d km/h\n", c.Marca, c.Modelo, c.Velocidade)
}

func (c Carro) ExibirVelocidade() {
    fmt.Printf("Velocidade atual do %s %s: %d km/h\n", c.Marca, c.Modelo, c.Velocidade)
}

func main() {
    carro1 := Carro{
        Marca:      "Toyota",
        Modelo:     "Corolla",
        Ano:        2020,
        Velocidade: 0,
    }

    carro2 := Carro{
        Marca:      "Honda",
        Modelo:     "Civic",
        Ano:        2018,
        Velocidade: 0,
    }

    carro3 := Carro{
        Marca:      "Ford",
        Modelo:     "Mustang",
        Ano:        2021,
        Velocidade: 0,
    }

    carro1.ExibirDetalhes()
    carro1.ExibirVelocidade()

    carro2.ExibirDetalhes()
    carro2.ExibirVelocidade()

    carro3.ExibirDetalhes()
    carro3.ExibirVelocidade()

    carro1.Acelerar(50)
    carro1.Frear(20)
    carro1.ExibirVelocidade()

    carro2.Acelerar(70)
    carro2.Frear(30)
    carro2.ExibirVelocidade()

    carro3.Acelerar(100)
    carro3.Frear(60)
    carro3.ExibirVelocidade()
}