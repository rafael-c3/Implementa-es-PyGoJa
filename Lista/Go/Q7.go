package main

import "fmt"

type Professor struct {
    Nome    string
    Escolas []*Escola
}

func (p Professor) ExibirEscolas() {
    fmt.Printf("Professor: %s\nLeciona nas seguintes escolas:\n", p.Nome)
    for _, escola := range p.Escolas {
        fmt.Printf("- %s\n", escola.Nome)
    }
}

func (p *Professor) AdicionarEscola(escola *Escola) {
    p.Escolas = append(p.Escolas, escola)
}

type Escola struct {
    Nome       string
    Professores []*Professor
}

func (e Escola) ExibirProfessores() {
    fmt.Printf("Escola: %s\nProfessores:\n", e.Nome)
    for _, professor := range e.Professores {
        fmt.Printf("- %s\n", professor.Nome)
    }
}

func (e *Escola) AdicionarProfessor(professor *Professor) {
    e.Professores = append(e.Professores, professor)
    professor.AdicionarEscola(e)
}

func main() {
    professor1 := &Professor{Nome: "Professor A"}
    professor2 := &Professor{Nome: "Professor B"}

    escola1 := &Escola{Nome: "Escola Alpha"}
    escola2 := &Escola{Nome: "Escola Beta"}

    escola1.AdicionarProfessor(professor1)
    escola1.AdicionarProfessor(professor2)

    escola2.AdicionarProfessor(professor1)

    escola1.ExibirProfessores()
    escola2.ExibirProfessores()

    professor1.ExibirEscolas()
    professor2.ExibirEscolas()
}
