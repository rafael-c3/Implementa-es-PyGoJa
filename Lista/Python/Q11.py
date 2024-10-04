from abc import ABC, abstractmethod

class Funcionario(ABC):
    
    @abstractmethod
    def calcularSalario(self):
        pass

class FuncionarioHorista(Funcionario):
    def __init__(self, nome, horas_trabalhadas, valor_por_hora):
        self.nome = nome
        self.horas_trabalhadas = horas_trabalhadas
        self.valor_por_hora = valor_por_hora

    def calcularSalario(self):
        return self.horas_trabalhadas * self.valor_por_hora

class FuncionarioAssalariado(Funcionario):
    def __init__(self, nome, salario_mensal):
        self.nome = nome
        self.salario_mensal = salario_mensal

    def calcularSalario(self):
        return self.salario_mensal

horista = FuncionarioHorista('João', 160, 25)
assalariado = FuncionarioAssalariado('Ana', 5000)

print(f'Salário de {horista.nome}: R$ {horista.calcularSalario():.2f}')
print(f'Salário de {assalariado.nome}: R$ {assalariado.calcularSalario():.2f}')