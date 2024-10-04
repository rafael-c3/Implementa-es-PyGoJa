package main

import "fmt"

// Definindo a classe (struct) Carro
type Carro struct {
    Marca  string
    Modelo string
    Ano    int
}

// Função para exibir os detalhes do carro
func (c Carro) ExibirDetalhes() {
    fmt.Printf("Marca: %s, Modelo: %s, Ano: %d\n", c.Marca, c.Modelo, c.Ano)
}

func main() {
    // Instanciando três objetos da classe Carro
    carro1 := Carro{
        Marca:  "Toyota",
        Modelo: "Corolla",
        Ano:    2020,
    }

    carro2 := Carro{
        Marca:  "Honda",
        Modelo: "Civic",
        Ano:    2018,
    }

    carro3 := Carro{
        Marca:  "Ford",
        Modelo: "Mustang",
        Ano:    2021,
    }

    // Exibindo os detalhes de cada carro
    carro1.ExibirDetalhes()
    carro2.ExibirDetalhes()
    carro3.ExibirDetalhes()
}
