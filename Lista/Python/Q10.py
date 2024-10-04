class Calculadora:
    def somar(self, *args):
        if len(args) < 2 or len(args) > 3:
            raise ValueError("Informe dois ou três números")
        return sum(args)

calc = Calculadora()
print(calc.somar(2, 3))
print(calc.somar(2, 3, 4))