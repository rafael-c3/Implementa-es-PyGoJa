package main

import "fmt"

type Animal interface {
    Som() string
}

type Cachorro struct {
    Nome string
}

func (c Cachorro) Som() string {
    return "Au Au"
}

type Gato struct {
    Nome string
}

func (g Gato) Som() string {
    return "Miau"
}

func FazerSom(animal Animal) {
    fmt.Println(animal.Som())
}

func main() {
    cachorro := Cachorro{Nome: "Rex"}
    gato := Gato{Nome: "Mimi"}

    fmt.Printf("O som do %s é: ", cachorro.Nome)
    FazerSom(cachorro)

    fmt.Printf("O som da %s é: ", gato.Nome)
    FazerSom(gato)
}
