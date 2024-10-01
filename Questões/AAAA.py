class Animal:
    def __init__(self, especie, nome):
        self.especie = especie
        self.nome = nome

        def emitir_som(self):
            pass

class Cachorro(Animal):
    def emitir_som(self):
        return "Au Au"
    
class Gato(Animal):
    def emitir_som(self):
        return "Miau"
    
def main():
    cachorro = Cachorro(especie="Caninca", nome="Snoopy")
    gato = Gato(especie="Felina", nome="Luna")

    print(f"{cachorro.nome}, especie {cachorro.especie} faz {cachorro.emitir_som()}")
    print(f"{gato.nome}, especie {gato.especie} faz {gato.emitir_som()}")

if __name__ == "__main__":
    main()