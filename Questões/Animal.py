class Animal:
    def __init__(self, nome):
        self.nome = nome

    def emitirsom(self):
        pass

class Mamífero(Animal):
    def __init__(self, nome):
        super().__init__(nome)

    def amamentar(self):
        print(f"{self.nome} está amamentando...")

class Ave(Animal):
    def __init__(self, nome):
        super().__init__(nome)

    def voar(self):
        print(f"{self.nome} está voando....")

class Morcego(Mamífero, Ave):
    def __init__(self, nome):
        super().__init__(nome)

    def emitir_som(self):
        print("Morcego fazendo ruidos noturnos")

def main():
    morcego = Morcego("Batemane")
    morcego.emitir_som()
    morcego.amamentar()
    morcego.voar()

if __name__ == '__main__':
    main()