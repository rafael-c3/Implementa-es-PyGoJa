class Motor {
    private String tipo;
    private int potencia;

    public Motor(String tipo, int potencia) {
        this.tipo = tipo;
        this.potencia = potencia;
    }

    public void exibirInformacoes() {
        System.out.println("Tipo do motor: " + tipo + ", Potência: " + potencia + " HP");
    }
}

class Carro {
    private String marca;
    private String modelo;
    private int ano;
    private Motor motor;

    public Carro(String marca, String modelo, int ano, Motor motor) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
        this.motor = motor;
    }

    public void exibirDetalhes() {
        System.out.println("Marca: " + marca + ", Modelo: " + modelo + ", Ano: " + ano);
        motor.exibirInformacoes();
    }
}

public class Main {
    public static void main(String[] args) {
        Motor motor1 = new Motor("Gasolina", 150);

        Carro carro1 = new Carro("Ford", "Mustang", 2022, motor1);

        carro1.exibirDetalhes();
    }
}
