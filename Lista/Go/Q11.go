package main

import "fmt"

type Funcionario interface {
    CalcularSalario() float64
}

type FuncionarioHorista struct {
    Nome       string
    HorasTrabalhadas float64
    ValorHora  float64
}

func (fh FuncionarioHorista) CalcularSalario() float64 {
    return fh.HorasTrabalhadas * fh.ValorHora
}

type FuncionarioAssalariado struct {
    Nome         string
    SalarioMensal float64
}

func (fa FuncionarioAssalariado) CalcularSalario() float64 {
    return fa.SalarioMensal
}

func ExibirSalario(f Funcionario) {
    fmt.Printf("Salário: R$%.2f\n", f.CalcularSalario())
}

func main() {
    funcionarioHorista := FuncionarioHorista{
        Nome:          "Alice",
        HorasTrabalhadas: 160,
        ValorHora:     50.0,
    }

    funcionarioAssalariado := FuncionarioAssalariado{
        Nome:         "Bob",
        SalarioMensal: 5000.0,
    }

    fmt.Println("Funcionário Horista:")
    ExibirSalario(funcionarioHorista)

    fmt.Println("Funcionário Assalariado:")
    ExibirSalario(funcionarioAssalariado)
}
