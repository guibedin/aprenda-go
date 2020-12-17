package main

import (
	"encoding/json"
	"fmt"
)

type pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func main() {

	p1 := pessoa{"Guilherme", 28}
	p2 := pessoa{"Camila", 27}

	fmt.Println(p1, p2)

	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("Error during marshal of %v: %v\n", p1, err)
	}
	fmt.Printf("%s\n", b)

	b, err = json.Marshal(p2)
	if err != nil {
		fmt.Printf("Error during marshal of %v: %v\n", p2, err)
	}
	fmt.Printf("%s\n", b)

}
