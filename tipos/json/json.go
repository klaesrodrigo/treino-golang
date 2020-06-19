package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float32  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := product{1, "Caneta", 4.20, []string{"Acessório", "Pincel"}}
	p1json, _ := json.Marshal(p1)
	fmt.Println(string(p1json))

	var p2 product
	jsonString := `{"id":2,"name": "Notbook", "price": 2566.99, "tags": ["Eletrnicos", "Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
}
