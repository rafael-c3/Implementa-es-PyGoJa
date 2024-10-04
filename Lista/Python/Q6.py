class Motor:
    def __init__(self, tipo, potencia):
        self.tipo = tipo
        self.potencia = potencia

    def exibir_detalhes(self):
        return f'Motor: {self.tipo}, Potência: {self.potencia} HP'

class Carro:
    def __init__(self, marca, modelo, ano, motor):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano
        self.motor = motor

    def exibir_detalhes(self):
        detalhes_carro = f'Marca: {self.marca}, Modelo: {self.modelo}, Ano: {self.ano}'
        detalhes_motor = self.motor.exibir_detalhes()
        return f'{detalhes_carro}, {detalhes_motor}'

motor_v6 = Motor('V6', 300)

carro = Carro('Ford', 'Mustang', 2021, motor_v6)

print(carro.exibir_detalhes())
