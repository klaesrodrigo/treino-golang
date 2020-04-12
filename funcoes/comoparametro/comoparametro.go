package main

import "fmt"

var mult = func(a, b int) int {
	return a * b
}

var exec = func(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	fmt.Println(exec(mult, 5, 3))
}
