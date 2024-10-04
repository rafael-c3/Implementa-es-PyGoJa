from abc import ABC, abstractmethod

class Imprimivel(ABC):
    
    @abstractmethod
    def imprimir(self):
        pass

class Relatorio(Imprimivel):
    def __init__(self, titulo, conteudo):
        self.titulo = titulo
        self.conteudo = conteudo

    def imprimir(self):
        return f'Relatório: {self.titulo}\nConteúdo: {self.conteudo}'

class Contrato(Imprimivel):
    def __init__(self, partes, termos):
        self.partes = partes
        self.termos = termos

    def imprimir(self):
        return f'Contrato entre: {self.partes}\nTermos: {self.termos}'

relatorio = Relatorio('Vendas Mensais', 'Relatório das vendas do mês de setembro')
contrato = Contrato('Empresa X e Empresa Y', 'Termos de prestação de serviços')

print(relatorio.imprimir())
print(contrato.imprimir())
