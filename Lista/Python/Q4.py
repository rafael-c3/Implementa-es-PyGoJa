class Animal:
    def __init__(self, nome):
        self.nome = nome

    def som(self):
        pass

class Cachorro(Animal):
    def som(self):
        return f'{self.nome} faz: Au Au!'

class Gato(Animal):
    def som(self):
        return f'{self.nome} faz: Miau!'

cachorro = Cachorro('Rex')
gato = Gato('Mingau')

print(cachorro.som())
print(gato.som())