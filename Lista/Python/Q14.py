class Configuracao:
    _instancia = None

    def __new__(cls):
        if cls._instancia is None:
            cls._instancia = super(Configuracao, cls).__new__(cls)
            cls._instancia._config = {}
        return cls._instancia

    def set_config(self, chave, valor):
        self._config[chave] = valor

    def get_config(self, chave):
        return self._config.get(chave, None)

config1 = Configuracao()
config1.set_config('host', 'localhost')

config2 = Configuracao()
config2.set_config('port', 8080)

print(config1 is config2)

print(config1.get_config('host'))
print(config2.get_config('port'))
print(config1.get_config('port'))
