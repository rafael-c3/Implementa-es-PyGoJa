package main

import "fmt"

type ContaBancaria struct {
    titular string
    saldo   float64
}

func NovaConta(titular string, saldoInicial float64) *ContaBancaria {
    return &ContaBancaria{
        titular: titular,
        saldo:   saldoInicial,
    }
}

func (c *ContaBancaria) Depositar(valor float64) {
    if valor > 0 {
        c.saldo += valor
        fmt.Printf("Depósito de R$%.2f realizado com sucesso. Saldo atual: R$%.2f\n", valor, c.saldo)
    } else {
        fmt.Println("Valor de depósito inválido.")
    }
}

func (c *ContaBancaria) Sacar(valor float64) {
    if valor > 0 && valor <= c.saldo {
        c.saldo -= valor
        fmt.Printf("Saque de R$%.2f realizado com sucesso. Saldo atual: R$%.2f\n", valor, c.saldo)
    } else if valor > c.saldo {
        fmt.Println("Saldo insuficiente para realizar o saque.")
    } else {
        fmt.Println("Valor de saque inválido.")
    }
}

func (c *ContaBancaria) ExibirSaldo() {
    fmt.Printf("Saldo atual de %s: R$%.2f\n", c.titular, c.saldo)
}

func main() {
    conta1 := NovaConta("Maria", 1000.00)
    conta2 := NovaConta("João", 500.00)

    conta1.ExibirSaldo()
    conta2.ExibirSaldo()

    conta1.Depositar(300.00)
    conta1.Sacar(200.00)
    conta1.ExibirSaldo()

    conta2.Depositar(100.00)
    conta2.Sacar(700.00)
    conta2.ExibirSaldo()
}
