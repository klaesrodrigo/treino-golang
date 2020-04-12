package main

import "fmt"

var media = func(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}

func main() {
	fmt.Println(media(5.0, 5.0, 5.0, 9.0))
	fmt.Println(media(8.0, 4.0, 8.6, 10.0))
	fmt.Println(media(8.0, 5.7, 4.0, 9.8))
}
