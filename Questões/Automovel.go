package main

import (
	"fmt"
)

// Motor struct
type Motor struct {
	Tipo    string
	Potencia int
}

// Ligar motor
func (m *Motor) Ligar() {
	fmt.Println("O motor está ligado")
}

// Desligar motor
func (m *Motor) Desligar() {
	fmt.Println("O motor está desligado")
}

// Pneu struct
type Pneu struct {
	Marca   string
	Tamanho int
}

// Inflar pneu
func (p *Pneu) Inflar() {
	fmt.Println("O pneu está inflado")
}

// Desinflar pneu
func (p *Pneu) Desinflar() {
	fmt.Println("O pneu está desinflado")
}

// Carro struct
type Carro struct {
	Marca   string
	Modelo  string
	Motor   *Motor
	Pneus   []*Pneu
}

// Ligar o carro
func (c *Carro) Ligar() {
	c.Motor.Ligar()
	fmt.Println("O carro está pronto para rodar")
}

// Desligar o carro
func (c *Carro) Desligar() {
	c.Motor.Desligar()
	fmt.Println("O carro foi desligado")
}

// TrocarPneu troca um pneu do carro
func (c *Carro) TrocarPneu(indice int, novoPneu *Pneu) {
	if indice >= 0 && indice < len(c.Pneus) {
		c.Pneus[indice] = novoPneu
		fmt.Printf("Pneu %d trocado\n", indice+1)
	} else {
		fmt.Println("Índice de pneu inválido")
	}
}

func main() {
	// Criar um carro
	carro := &Carro{
		Marca:  "Toyota",
		Modelo: "Corolla",
		Motor:  &Motor{"Gasolina", 150},
		Pneus:  []*Pneu{
			&Pneu{"Pirelli", 18},
			&Pneu{"Pirelli", 18},
			&Pneu{"Pirelli", 18},
			&Pneu{"Pirelli", 18},
		},
	}

	carro.Ligar()

	// Trocar pneu
	novoPneu := &Pneu{"Michelin", 27}
	carro.TrocarPneu(2, novoPneu)

	carro.Desligar()
}
