package main

import (
    "fmt"
)

type SaldoInsuficienteException struct {
    SaldoAtual float64
    ValorSaque float64
}

func (e *SaldoInsuficienteException) Error() string {
    return fmt.Sprintf("Saldo insuficiente: saldo atual é R$%.2f, tentativa de saque é R$%.2f", e.SaldoAtual, e.ValorSaque)
}

type ContaBancaria struct {
    Titular string
    saldo   float64
}

func (c *ContaBancaria) Depositar(valor float64) {
    c.saldo += valor
}

func (c *ContaBancaria) Sacar(valor float64) error {
    if valor > c.saldo {
        return &SaldoInsuficienteException{SaldoAtual: c.saldo, ValorSaque: valor}
    }
    c.saldo -= valor
    return nil
}

func (c *ContaBancaria) ObterSaldo() float64 {
    return c.saldo
}

func main() {
    conta := ContaBancaria{Titular: "Alice", saldo: 100.0}

    err := conta.Sacar(150.0)
    if err != nil {
        fmt.Println("Erro:", err)
    } else {
        fmt.Println("Saque realizado com sucesso!")
    }

    fmt.Printf("Saldo atual: R$%.2f\n", conta.ObterSaldo())
}
