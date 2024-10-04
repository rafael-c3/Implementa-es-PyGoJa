class Empregado:
    def __init__(self, nome, cargo, salario):
        self.nome = nome
        self.cargo = cargo
        self.salario = salario

    def exibir_detalhes(self):
        return f'Nome: {self.nome}, Cargo: {self.cargo}, Salário: R${self.salario:.2f}'

class Empresa:
    def __init__(self, nome_empresa):
        self.nome_empresa = nome_empresa
        self.empregados = []

    def adicionar_empregado(self, empregado):
        self.empregados.append(empregado)

    def exibir_empregados(self):
        print(f'Empregados da {self.nome_empresa}:')
        for empregado in self.empregados:
            print(empregado.exibir_detalhes())

empregado1 = Empregado('João', 'Desenvolvedor', 5000)
empregado2 = Empregado('Ana', 'Gerente de Projetos', 8000)
empregado3 = Empregado('Pedro', 'Designer', 4000)

empresa = Empresa('Tech Solutions')

empresa.adicionar_empregado(empregado1)
empresa.adicionar_empregado(empregado2)
empresa.adicionar_empregado(empregado3)

empresa.exibir_empregados()