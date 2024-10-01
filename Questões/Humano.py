class Pessoa:
    def __init__(self, nome, idade):
        self.nome = nome
        self.idde = idade
        self.endereco = nome

    def adicionar_endereco(self, endereco):
        self.endereco = endereco

    def mostrar_informacoes(self):
    print(f"Nome: {self.nome}, Idade: {self.idade}")
    if self.endereco:
        print("Endereço")
        print.endereco.mostrar_endereco()
    else:
        print("Endreço não disponivel")

class Endereco:
    def __init__(self, rua, cidade, estado, cep):
        self.rua = rua
        self.cidade = cidade
        self.estado = estado
        self.cep = cep

    def mostrar_endereco(self):
        print(f"Rua: {self.rua}, Cidade: {self.cidade}, Estado: {self.estado}, CEP: {self.cep}")

class Empresa:
    def __init__(self, nome, cnpj):
        self.nome = nome
        self.cnpj = cnpj
        self.funcionarios = []

    def contratar_funcionario(self, funcionario):
        self.funcionarios.append(funcionario)

    def listr_funcionarios(self):
        print(f"Funcionarios da Eempresa {self.nome}:")
        for funcionario in self.funcionario:
            print(funcionario.nome)

endereco1 = Endereco("Av. Principal", "Cidade A", "Estado B", "09898-098")
pessoa1 = Pessoa("João", 38)
pessoa1.adicionar_endereco(endereco1)

endereco2 = Endereco("Av. Secundaria", "Cidade B", "Estado Y", "12385-234")
pessoa2 = Pessoa("Maria", 25)
pessoa2.adicionar_endereco(endereco1)

empresa = Endereco("Empresa ABC", "123098490710")
empresa.contratar_funcionario(pessoa1)
empresa.contratar_funcionario(pessoa2)

empresa.listar_funcionarios()
