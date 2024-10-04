class ContaBancaria {
    private double saldo;
    private String titular;

    public ContaBancaria(String titular) {
        this.titular = titular;
        this.saldo = 0.0;
    }

    public void depositar(double valor) {
        if (valor > 0) {
            saldo += valor;
            System.out.println("Depósito de " + valor + " realizado. Saldo atual: " + saldo);
        } else {
            System.out.println("Valor de depósito deve ser positivo.");
        }
    }

    public void sacar(double valor) {
        if (valor > 0) {
            if (valor <= saldo) {
                saldo -= valor;
                System.out.println("Saque de " + valor + " realizado. Saldo atual: " + saldo);
            } else {
                System.out.println("Saldo insuficiente para saque de " + valor + ". Saldo atual: " + saldo);
            }
        } else {
            System.out.println("Valor de saque deve ser positivo.");
        }
    }

    public void exibirSaldo() {
        System.out.println("Saldo atual de " + titular + ": " + saldo);
    }

    public String getTitular() {
        return titular;
    }
}

public class Main {
    public static void main(String[] args) {
        ContaBancaria conta1 = new ContaBancaria("João Silva");

        conta1.exibirSaldo();

        conta1.depositar(500.00);
        conta1.sacar(200.00);
        conta1.sacar(350.00);
        conta1.depositar(-50.00);
        conta1.exibirSaldo();
    }
}
