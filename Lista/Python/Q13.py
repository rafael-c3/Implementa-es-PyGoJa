class Matematica:
    @staticmethod
    def fatorial(n):
        if n < 0:
            raise ValueError("O fatorial não está definido para números negativos.")
        elif n == 0 or n == 1:
            return 1
        else:
            resultado = 1
            for i in range(2, n + 1):
                resultado *= i
            return resultado

numero = 5
print(f'O fatorial de {numero} é: {Matematica.fatorial(numero)}')

numero_negativo = -3
try:
    print(f'O fatorial de {numero_negativo} é: {Matematica.fatorial(numero_negativo)}')
except ValueError as e:
    print(e)