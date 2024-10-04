package main

import (
    "fmt"
)

type Produto struct {
    Nome  string
    Preco float64
}

func (p1 Produto) Somar(p2 Produto) Produto {
    return Produto{
        Nome:  p1.Nome + " e " + p2.Nome,
        Preco: p1.Preco + p2.Preco,
    }
}

type Matematica struct{}

func (Matematica) Fatorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * Matematica{}.Fatorial(n-1)
}

func (Matematica) Somar(a, b int) int {
    return a + b
}

func main() {
    produto1 := Produto{Nome: "Produto A", Preco: 10.50}
    produto2 := Produto{Nome: "Produto B", Preco: 15.75}

    produtoSoma := produto1.Somar(produto2)
    fmt.Printf("Produto combinado: %s, Preço total: R$%.2f\n", produtoSoma.Nome, produtoSoma.Preco)

    matematica := Matematica{}

    num := 5
    fatorial := matematica.Fatorial(num)
    fmt.Printf("Fatorial de %d: %d\n", num, fatorial)

    soma := matematica.Somar(3, 7)
    fmt.Printf("Soma de 3 e 7: %d\n", soma)
}
