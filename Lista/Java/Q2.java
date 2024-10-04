class Carro {
    private String marca;
    private String modelo;
    private int ano;
    private int velocidade;

    public Carro(String marca, String modelo, int ano) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
        this.velocidade = 0;
    }

    public void exibirDetalhes() {
        System.out.println("Marca: " + marca + ", Modelo: " + modelo + ", Ano: " + ano);
        exibirVelocidade();
    }

    public void acelerar(int incremento) {
        if (incremento > 0) {
            velocidade += incremento;
            System.out.println("Acelerando. Velocidade atual: " + velocidade + " km/h");
        } else {
            System.out.println("Incremento deve ser positivo.");
        }
    }

    public void frear(int decremento) {
        if (decremento > 0) {
            velocidade -= decremento;
            if (velocidade < 0) {
                velocidade = 0;
            }
            System.out.println("Freando. Velocidade atual: " + velocidade + " km/h");
        } else {
            System.out.println("Decremento deve ser positivo.");
        }
    }

    public void exibirVelocidade() {
        System.out.println("Velocidade atual: " + velocidade + " km/h");
    }
}

public class Main {
    public static void main(String[] args) {
        Carro carro1 = new Carro("Ford", "Mustang", 2022);
        Carro carro2 = new Carro("Chevrolet", "Camaro", 2021);
        Carro carro3 = new Carro("Toyota", "Corolla", 2023);

        carro1.exibirDetalhes();
        carro2.exibirDetalhes();
        carro3.exibirDetalhes();

        carro1.acelerar(50); // Acelera o carro 1
        carro1.frear(20);    // Freia o carro 1
        carro2.acelerar(30); // Acelera o carro 2
        carro3.frear(10);    // Freia o carro 3
        carro3.acelerar(70); // Acelera o carro 3
    }
}
