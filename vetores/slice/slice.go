package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [6]int{1, 2, 3, 4, 5, 6}
	s1 := []int{1, 2, 3}

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{5, 6, 7, 8, 9}

	s2 := a2[0:2]
	fmt.Println(a2, s2)

	s3 := a2[1:4]
	fmt.Println(a2, s3)

	s4 := s3[1:2]
	fmt.Println(s3, s4)
}
