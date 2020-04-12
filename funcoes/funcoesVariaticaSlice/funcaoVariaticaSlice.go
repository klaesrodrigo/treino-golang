package main

import "fmt"

func printAprovados(aprovados ...string) {
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {

	aprovados := []string{"Rogério", "Fábio", "Rafa", "Renata", "Bianca", "Marcelo", "Paulo"}
	fmt.Println("Lista de aprovados")
	printAprovados(aprovados...)
}
