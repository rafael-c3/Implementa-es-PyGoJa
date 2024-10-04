class Matematica {

    public static long fatorial(int numero) {
        if (numero < 0) {
            throw new IllegalArgumentException("O número deve ser não-negativo.");
        }
        long resultado = 1;
        for (int i = 1; i <= numero; i++) {
            resultado *= i;
        }
        return resultado;
    }
}

public class Main {
    public static void main(String[] args) {
        int numero = 5;
        long resultado = Matematica.fatorial(numero);
        System.out.println("Fatorial de " + numero + " é: " + resultado);

        numero = 10;
        resultado = Matematica.fatorial(numero);
        System.out.println("Fatorial de " + numero + " é: " + resultado);

        try {
            numero = -3;
            resultado = Matematica.fatorial(numero);
            System.out.println("Fatorial de " + numero + " é: " + resultado);
        } catch (IllegalArgumentException e) {
            System.out.println(e.getMessage());
        }
    }
}
