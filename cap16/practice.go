package main

import (
	"golang.org/x/crypto/bcrypt"

	"encoding/json"
	"fmt"
	"sort"
)

type pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

type porIdade []pessoa

// Metodos para implementar a interface Interface do pacote sort
func (pi porIdade) Len() int           { return len(pi) }
func (pi porIdade) Less(i, j int) bool { return pi[i].Idade < pi[j].Idade }
func (pi porIdade) Swap(i, j int)      { pi[i], pi[j] = pi[j], pi[i] }

func testandoMarshal() {
	fmt.Println("Marshal")

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
	fmt.Printf("%s\n\n", b)
}

func testandoUnmarshal() {
	fmt.Println("Unmarshal")

	var sb []byte
	sb = []byte(`{"nome":"Guilherme","idade":28}`)
	fmt.Println(string(sb))

	var p pessoa
	err := json.Unmarshal(sb, &p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p)
	fmt.Println()
}

func testandoSort() {
	fmt.Println("Sort padrão")

	si := []int{1, 5, 4, 3, 2, 10, 9}
	fmt.Println(si)

	sort.Ints(si)
	fmt.Println(si)

	fmt.Println()

	fmt.Println("Sort pessoa por idade")

	pessoas := []pessoa{pessoa{"Guilherme", 28}, pessoa{"Camila", 27}}

	fmt.Println(pessoas)
	sort.Sort(porIdade(pessoas))
	fmt.Println(pessoas)
}

func testandoBcrypt() {
	senha := "abcabc"

	hash, err := bcrypt.GenerateFromPassword([]byte(senha), 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(hash))

	err = bcrypt.CompareHashAndPassword(hash, []byte(senha))
	if err != nil {
		fmt.Println("Senha incorreta!", err)
	} else {
		fmt.Println("Senha correta!")
	}
}

func main() {

	//testandoMarshal()
	//testandoUnmarshal()
	//testandoSort()
	testandoBcrypt()

}
