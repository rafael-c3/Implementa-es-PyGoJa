package main

import "fmt"

type Imprimivel interface {
    Imprimir()
}

type Relatorio struct {
    Titulo  string
    Conteudo string
}

func (r Relatorio) Imprimir() {
    fmt.Printf("Relatório: %s\nConteúdo: %s\n", r.Titulo, r.Conteudo)
}

type Contrato struct {
    NomeEmpresa string
    Detalhes    string
}

func (c Contrato) Imprimir() {
    fmt.Printf("Contrato com: %s\nDetalhes: %s\n", c.NomeEmpresa, c.Detalhes)
}

func ProcessarImpressao(i Imprimivel) {
    i.Imprimir()
}

func main() {
    relatorio := Relatorio{
        Titulo:   "Relatório Anual de Vendas",
        Conteudo: "Resumo das vendas do ano de 2023.",
    }

    contrato := Contrato{
        NomeEmpresa: "Empresa XYZ",
        Detalhes:    "Contrato de fornecimento de tecnologia.",
    }

    ProcessarImpressao(relatorio)
    ProcessarImpressao(contrato)
}
