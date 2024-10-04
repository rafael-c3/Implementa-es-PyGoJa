package main

import "fmt"

type Empregado struct {
    Nome   string
    Cargo  string
    Salario float64
}

func (e Empregado) ExibirDetalhes() {
    fmt.Printf("Nome: %s, Cargo: %s, Salário: R$%.2f\n", e.Nome, e.Cargo, e.Salario)
}

type Empresa struct {
    Nome       string
    Empregados []Empregado
}

func (e *Empresa) AdicionarEmpregado(empregado Empregado) {
    e.Empregados = append(e.Empregados, empregado)
}

func (e Empresa) ExibirEmpregados() {
    fmt.Printf("Empresa: %s\n", e.Nome)
    fmt.Println("Lista de empregados:")
    for _, empregado := range e.Empregados {
        empregado.ExibirDetalhes()
    }
}

func (e Empresa) CalcularTotalSalarios() float64 {
    total := 0.0
    for _, empregado := range e.Empregados {
        total += empregado.Salario
    }
    return total
}

func main() {
    empresa := Empresa{
        Nome: "Tech Solutions",
    }

    empregado1 := Empregado{Nome: "Alice", Cargo: "Desenvolvedora", Salario: 5000.00}
    empregado2 := Empregado{Nome: "Bob", Cargo: "Analista de Sistemas", Salario: 4500.00}
    empregado3 := Empregado{Nome: "Carlos", Cargo: "Gerente de Projetos", Salario: 7000.00}

    empresa.AdicionarEmpregado(empregado1)
    empresa.AdicionarEmpregado(empregado2)
    empresa.AdicionarEmpregado(empregado3)

    empresa.ExibirEmpregados()

    totalSalarios := empresa.CalcularTotalSalarios()
    fmt.Printf("Total de salários pagos pela empresa: R$%.2f\n", totalSalarios)
}
