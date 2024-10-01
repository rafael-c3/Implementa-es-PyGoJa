package main

import (
	"fmt"
)

// Pessoa struct
type Pessoa struct {
	Nome     string
	Idade    int
	Endereco *Endereco
}

// AdicionarEndereco adiciona um endereço à pessoa
func (p *Pessoa) AdicionarEndereco(endereco *Endereco) {
	p.Endereco = endereco
}

// MostrarInformacoes exibe as informações da pessoa
func (p *Pessoa) MostrarInformacoes() {
	fmt.Printf("Nome: %s, Idade: %d\n", p.Nome, p.Idade)
	if p.Endereco != nil {
		fmt.Println("Endereço:")
		p.Endereco.MostrarEndereco()
	} else {
		fmt.Println("Endereço não disponível")
	}
}

// Endereco struct
type Endereco struct {
	Rua    string
	Cidade string
	Estado string
	CEP    string
}

// MostrarEndereco exibe o endereço
func (e *Endereco) MostrarEndereco() {
	fmt.Printf("Rua: %s, Cidade: %s, Estado: %s, CEP: %s\n", e.Rua, e.Cidade, e.Estado, e.CEP)
}

// Empresa struct
type Empresa struct {
	Nome        string
	CNPJ        string
	Funcionarios []*Pessoa
}

// ContratarFuncionario adiciona um funcionário à empresa
func (e *Empresa) ContratarFuncionario(funcionario *Pessoa) {
	e.Funcionarios = append(e.Funcionarios, funcionario)
}

// ListarFuncionarios exibe a lista de funcionários da empresa
func (e *Empresa) ListarFuncionarios() {
	fmt.Printf("Funcionários da Empresa %s:\n", e.Nome)
	for _, funcionario := range e.Funcionarios {
		fmt.Println(funcionario.Nome)
	}
}

func main() {
	endereco1 := &Endereco{"Av. Principal", "Cidade A", "Estado B", "09898-098"}
	pessoa1 := &Pessoa{"João", 38, nil}
	pessoa1.AdicionarEndereco(endereco1)

	endereco2 := &Endereco{"Av. Secundária", "Cidade B", "Estado Y", "12385-234"}
	pessoa2 := &Pessoa{"Maria", 25, nil}
	pessoa2.AdicionarEndereco(endereco2)

	empresa := &Empresa{"Empresa ABC", "123098490710", nil}
	empresa.ContratarFuncionario(pessoa1)
	empresa.ContratarFuncionario(pessoa2)

	empresa.ListarFuncionarios()
}
