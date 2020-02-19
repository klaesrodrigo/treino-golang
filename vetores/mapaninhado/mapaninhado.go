package main

import "fmt"

func main() {
	salarios := map[string]map[string]float64{
		"A": {
			"Alberto Rosa":  14256.35,
			"Amanda Baldes": 4256.35,
		},
		"M": {
			"Mariana Silva": 1500.0,
		},
		"R": {
			"Ricardo Valdes": 9252.65,
			"Rodrigo Klaes":  11252.65,
		},
	}

	fmt.Println(salarios)

	for nome, salario := range salarios {
		fmt.Println(nome, salario)
	}

	delete(salarios, "R")

	for nome, salario := range salarios {
		fmt.Println(nome, salario)
	}
}
