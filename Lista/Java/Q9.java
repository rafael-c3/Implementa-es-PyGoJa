interface Imprimivel {
    void imprimir();
}

class Relatorio implements Imprimivel {
    private String titulo;
    private String conteudo;

    public Relatorio(String titulo, String conteudo) {
        this.titulo = titulo;
        this.conteudo = conteudo;
    }

    @Override
    public void imprimir() {
        System.out.println("----- Relatório -----");
        System.out.println("Título: " + titulo);
        System.out.println("Conteúdo: " + conteudo);
        System.out.println("---------------------");
    }
}

class Contrato implements Imprimivel {
    private String parteA;
    private String parteB;

    public Contrato(String parteA, String parteB) {
        this.parteA = parteA;
        this.parteB = parteB;
    }

    @Override
    public void imprimir() {
        System.out.println("----- Contrato -----");
        System.out.println("Parte A: " + parteA);
        System.out.println("Parte B: " + parteB);
        System.out.println("---------------------");
    }
}

public class Main {
    public static void main(String[] args) {
        Imprimivel relatorio = new Relatorio("Relatório Mensal", "Este é o conteúdo do relatório mensal.");
        Imprimivel contrato = new Contrato("Empresa XYZ", "Serviço de Consultoria");

        relatorio.imprimir();
        System.out.println();
        contrato.imprimir();
    }
}
