import java.util.ArrayList;
import java.util.List;

class Empregado {
    private String nome;
    private String cargo;
    private double salario;

    public Empregado(String nome, String cargo, double salario) {
        this.nome = nome;
        this.cargo = cargo;
        this.salario = salario;
    }

    public void exibirInformacoes() {
        System.out.println("Nome: " + nome + ", Cargo: " + cargo + ", Salário: R$ " + salario);
    }
    public String getNome() {
        return nome;
    }

    public String getCargo() {
        return cargo;
    }

    public double getSalario() {
        return salario;
    }
}

class Empresa {
    private String nome;
    private List<Empregado> empregados;

    public Empresa(String nome) {
        this.nome = nome;
        this.empregados = new ArrayList<>();
    }

    public void adicionarEmpregado(Empregado empregado) {
        empregados.add(empregado);
    }

    public void exibirInformacoes() {
        System.out.println("Empresa: " + nome);
        System.out.println("Empregados:");
        for (Empregado empregado : empregados) {
            empregado.exibirInformacoes();
        }
    }
}

public class Main {
    public static void main(String[] args) {
        Empresa empresa = new Empresa("Tech Solutions");

        Empregado empregado1 = new Empregado("João", "Desenvolvedor", 5000.00);
        Empregado empregado2 = new Empregado("Maria", "Gerente de Projetos", 8000.00);
        Empregado empregado3 = new Empregado("Carlos", "Designer", 4000.00);

        empresa.adicionarEmpregado(empregado1);
        empresa.adicionarEmpregado(empregado2);
        empresa.adicionarEmpregado(empregado3);

        empresa.exibirInformacoes();
    }
}
