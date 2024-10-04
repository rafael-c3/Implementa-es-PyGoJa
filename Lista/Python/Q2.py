class Carro:
    def __init__(self, marca, modelo, ano):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano
        self.velocidade = 0 

    def exibir_detalhes(self):
        print(f'Marca: {self.marca}, Modelo: {self.modelo}, Ano: {self.ano}')

    def acelerar(self, incremento):
        self.velocidade += incremento
        print(f'{self.modelo} acelerou para {self.velocidade} km/h')

    def frear(self, decremento):
        self.velocidade -= decremento
        if self.velocidade < 0:
            self.velocidade = 0
        print(f'{self.modelo} reduziu para {self.velocidade} km/h')

    def exibir_velocidade(self):
        print(f'Velocidade atual do {self.modelo}: {self.velocidade} km/h')

carro1 = Carro('Toyota', 'Corolla', 2020)
carro2 = Carro('Honda', 'Civic', 2018)
carro3 = Carro('Ford', 'Mustang', 2021)

carro1.exibir_detalhes()
carro1.acelerar(30)
carro1.frear(10)
carro1.exibir_velocidade()

carro2.exibir_detalhes()
carro2.acelerar(50)
carro2.frear(30)
carro2.exibir_velocidade()

carro3.exibir_detalhes()
carro3.acelerar(100)
carro3.frear(50)
carro3.exibir_velocidade()
