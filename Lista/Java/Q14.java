class Configuracao {
    private static Configuracao instancia;

    private String ambiente;
    private String bancoDeDados;

    private Configuracao() {
        this.ambiente = "Desenvolvimento";
        this.bancoDeDados = "MySQL";
    }

    public static Configuracao getInstancia() {
        if (instancia == null) {
            instancia = new Configuracao();
        }
        return instancia;
    }

    public String getAmbiente() {
        return ambiente;
    }

    public void setAmbiente(String ambiente) {
        this.ambiente = ambiente;
    }

    public String getBancoDeDados() {
        return bancoDeDados;
    }

    public void setBancoDeDados(String bancoDeDados) {
        this.bancoDeDados = bancoDeDados;
    }
}

public class Main {
    public static void main(String[] args) {
        Configuracao config1 = Configuracao.getInstancia();
        System.out.println("Ambiente: " + config1.getAmbiente());
        System.out.println("Banco de Dados: " + config1.getBancoDeDados());

        config1.setAmbiente("Produção");
        config1.setBancoDeDados("PostgreSQL");

        Configuracao config2 = Configuracao.getInstancia();
        System.out.println("\nApós modificações:");
        System.out.println("Ambiente: " + config2.getAmbiente());
        System.out.println("Banco de Dados: " + config2.getBancoDeDados());

        System.out.println("\nAs instâncias são iguais? " + (config1 == config2));
    }
}
