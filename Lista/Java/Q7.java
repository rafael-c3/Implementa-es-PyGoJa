import java.util.ArrayList;
import java.util.List;

class Professor {
    private String nome;
    private List<Escola> escolas;

    public Professor(String nome) {
        this.nome = nome;
        this.escolas = new ArrayList<>();
    }

    public void adicionarEscola(Escola escola) {
        if (!escolas.contains(escola)) {
            escolas.add(escola);
            escola.adicionarProfessor(this);
        }
    }

    public void exibirInformacoes() {
        System.out.println("Professor: " + nome);
        System.out.println("Escolas onde leciona: ");
        for (Escola escola : escolas) {
            System.out.println(" - " + escola.getNome());
        }
    }
}

class Escola {
    private String nome;
    private List<Professor> professores;

    public Escola(String nome) {
        this.nome = nome;
        this.professores = new ArrayList<>();
    }

    public void adicionarProfessor(Professor professor) {
        if (!professores.contains(professor)) {
            professores.add(professor);
        }
    }

    public String getNome() {
        return nome;
    }

    public void exibirInformacoes() {
        System.out.println("Escola: " + nome);
        System.out.println("Professores: ");
        for (Professor professor : professores) {
            System.out.println(" - " + professor.nome);
        }
    }
}

public class Main {
    public static void main(String[] args) {
        Escola escola1 = new Escola("Escola A");
        Escola escola2 = new Escola("Escola B");

        Professor professor1 = new Professor("Carlos");
        Professor professor2 = new Professor("Maria");

        professor1.adicionarEscola(escola1);
        professor1.adicionarEscola(escola2);
        professor2.adicionarEscola(escola1);

        professor1.exibirInformacoes();
        System.out.println();
        professor2.exibirInformacoes();
        System.out.println();

        escola1.exibirInformacoes();
        System.out.println();
        escola2.exibirInformacoes();
    }
}
