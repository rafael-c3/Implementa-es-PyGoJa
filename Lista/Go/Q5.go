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

func FazerSons(animais []Animal) {
    for _, animal := range animais {
        fmt.Println(animal.Som())
    }
}

func main() {
    cachorro := Cachorro{Nome: "Rex"}
    gato := Gato{Nome: "Mimi"}

    animais := []Animal{cachorro, gato}

    FazerSons(animais)
}
