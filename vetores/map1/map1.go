package main

import "fmt"

func main() {
	//maps devem ser inicializados
	m1 := make(map[int]string)

	m1[12345678998] = "Rodrigo"
	m1[98765432145] = "Alberto"
	m1[45678912358] = "Lucas"

	for cpf, nome := range m1 {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(m1[12345678998])
	delete(m1, 12345678998)
	fmt.Println(m1)
}
