class ContaBancaria:
    def __init__(self, titular, saldo_inicial=0):
        self.__saldo = saldo_inicial 
        self.titular = titular

    def depositar(self, valor):
        if valor > 0:
            self.__saldo += valor
            print(f'Depósito de {valor} realizado com sucesso!')
        else:
            print('Valor de depósito inválido!')

    def sacar(self, valor):
        if 0 < valor <= self.__saldo:
            self.__saldo -= valor
            print(f'Saque de {valor} realizado com sucesso!')
        else:
            print('Saque inválido ou saldo insuficiente!')

    def exibir_saldo(self):
        return self.__saldo

conta = ContaBancaria('João', 500)
conta.depositar(200)
conta.sacar(100)
print(f'Saldo atual: {conta.exibir_saldo()}')
