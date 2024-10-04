class Calculadora {

    public double somar(double a, double b) {
        return a + b;
    }

    public double somar(double a, double b, double c) {
        return a + b + c;
    }
}

public class Main {
    public static void main(String[] args) {
        Calculadora calculadora = new Calculadora();

        double resultado1 = calculadora.somar(5.0, 3.0);
        System.out.println("Resultado da soma de 5.0 e 3.0: " + resultado1);

        double resultado2 = calculadora.somar(5.0, 3.0, 2.0);
        System.out.println("Resultado da soma de 5.0, 3.0 e 2.0: " + resultado2);
    }
}
