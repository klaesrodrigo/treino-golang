package main

import "fmt"

type imprimivel interface {
	toString() string
}

type Pessoa struct {
	nome      string
	sobrenome string
}

type Produto struct {
	nome  string
	preco float64
}

func (p Pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p Produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = Pessoa{"Rodrigo", "Klaes"}
	fmt.Println(coisa.toString())
	imprimir(coisa)
	coisa = Produto{"Caneta", 1.30}
	fmt.Println(coisa.toString())
	imprimir(coisa)

}
