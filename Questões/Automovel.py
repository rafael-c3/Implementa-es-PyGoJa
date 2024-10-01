class Motor:
    def __init__(self, tipo, potencia):
        self.tipo = tipo
        self.potencia = potencia

    def ligar(self):
        print("O motor está ligaodo")

    def desligar(self):
        print("O motor está desligado")

class Pneu:
    def __init__(self, marca, tamanho):
        self.marca = marca
        self.tamanho = tamanho

    def inflar(self):
        print("O pneu está inflado")

    def desinflar(self):
        print("O pneu está deseinflado")

class Carro:
    def __init__ (self, marca, modelo):
        self.marca = marca
        self.modelo = modelo
        self.motor = Motor("Gasolina", 150)
        self.pneus = [Pneu("Pirelli", 18) for _ in range(4)]

    def ligar(self):
        self.motor.ligar()
        print("O carro está pronto para rodar")

    def desligar(self):
        self.motor.desligar()
        print("O carro foi desligado")

    def trocar_pneu(self, indice, novo_pneu):
        self.pneus[indice] = novo_pneu
        print(f"Pneu {indice +1} trocado")

carro = Carro("Toyota", "Corolla")
carro.ligar()

novo_pneu = Pneu("Michelin", 27)
carro.trocar_pneu(2, novo_pneu)
carro.desligar()