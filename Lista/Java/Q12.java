class Produto {
    private String nome;
    private double preco;

    public Produto(String nome, double preco) {
        this.nome = nome;
        this.preco = preco;
    }

    public double getPreco() {
        return preco;
    }

    public Produto somar(Produto outroProduto) {
        double novoPreco = this.preco + outroProduto.getPreco();
        return new Produto("Combo de " + this.nome + " e " + outroProduto.nome, novoPreco);
    }

    public void exibirInformacoes() {
        System.out.println("Produto: " + nome + ", Preço: R$ " + preco);
    }
}

public class Main {
    public static void main(String[] args) {
        Produto produto1 = new Produto("Produto A", 50.00);
        Produto produto2 = new Produto("Produto B", 30.00);

        produto1.exibirInformacoes();
        produto2.exibirInformacoes();

        Produto combo = produto1.somar(produto2);

        combo.exibirInformacoes();
    }
}
