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

def emitir_sons(animais):
    for animal in animais:
        print(animal.som())

cachorro = Cachorro('Rex')
gato = Gato('Mingau')

animais = [cachorro, gato]

emitir_sons(animais)
