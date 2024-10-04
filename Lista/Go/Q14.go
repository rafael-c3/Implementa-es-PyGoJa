package main

import (
    "fmt"
    "sync"
)

type Configuracao struct {
    Ambiente string
    Versao   string
}

var instance *Configuracao

var mu sync.Mutex

func GetConfiguracao() *Configuracao {
    mu.Lock()
    defer mu.Unlock()

    if instance == nil {
        instance = &Configuracao{
            Ambiente: "Desenvolvimento",
            Versao:   "1.0.0",
        }
    }
    return instance
}

func main() {
    config1 := GetConfiguracao()
    fmt.Printf("Configuração 1: Ambiente=%s, Versão=%s\n", config1.Ambiente, config1.Versao)

    config2 := GetConfiguracao()
    fmt.Printf("Configuração 2: Ambiente=%s, Versão=%s\n", config2.Ambiente, config2.Versao)

    if config1 == config2 {
        fmt.Println("As duas instâncias são as mesmas.")
    }
}
