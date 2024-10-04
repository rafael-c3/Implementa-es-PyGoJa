class SaldoInsuficienteException(Exception):
    def __init__(self, saldo_atual, valor_saque):
        super().__init__(f'Saldo insuficiente: R$ {saldo_atual:.2f}. Tentativa de saque: R$ {valor_saque:.2f}.')
        self.saldo_atual = saldo_atual
        self.valor_saque = valor_saque

class ContaBancaria:
    def __init__(self, titular, saldo_inicial=0):
        self.titular = titular
        self.saldo = saldo_inicial

    def depositar(self, valor):
        if valor > 0:
            self.saldo += valor
        else:
            raise ValueError("O valor do depósito deve ser positivo.")

    def sacar(self, valor):
        if valor > self.saldo:
            raise SaldoInsuficienteException(self.saldo, valor)
        self.saldo -= valor

    def obter_saldo(self):
        return self.saldo

conta = ContaBancaria('João', 1000)

try:
    conta.sacar(1200)
except SaldoInsuficienteException as e:
    print(e) 

print(f'Saldo atual: R$ {conta.obter_saldo():.2f}')
