package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto

	produto1 = produto{
		nome:     "Tesoura",
		preco:    10.0,
		desconto: 0.1,
	}

	fmt.Println(produto1.precoComDesconto())

	produto2 := produto{"Caneta", 5, 0.05}
	fmt.Println(produto2.precoComDesconto())
}
