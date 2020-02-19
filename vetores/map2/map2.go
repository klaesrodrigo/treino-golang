package main

import "fmt"

func main() {
	salarios := map[string]float64{
		"Rodrigo Klaes": 11252.65,
		"Aberto Rosa":   14256.35,
		"Mariana Silva": 1500.0,
	}

	fmt.Println(salarios)

	salarios["Adriano Imperador"] = 54223.66

	delete(salarios, "inexistente")

	for nome, salario := range salarios {
		fmt.Printf("%s - %f\n", nome, salario)
	}
}
