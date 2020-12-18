package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

/* EX 1
- Partindo do código abaixo, utilize marshal para transformar []user em JSON.
    - https://play.golang.org/p/U0jea43X55​
- Atenção! Tem pegadinha aqui. */
type user struct {
	First string
	Age   int
}

func ex1() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}
	fmt.Println(users)

	uj, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(uj))
}

/* EX 2
- Partindo do código abaixo, utilize unmarshal e demonstre os valores.
	- https://play.golang.org/p/b_UuCcZag9 */
type character struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// EX 5 - implementando a interface Interface para o sort
type byAge []character

func (x byAge) Len() int           { return len(x) }
func (x byAge) Less(i, j int) bool { return x[i].Age < x[j].Age }
func (x byAge) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byLast []character

func (x byLast) Len() int           { return len(x) }
func (x byLast) Less(i, j int) bool { return x[i].Last < x[j].Last }
func (x byLast) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func ex2() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	var c []character
	err := json.Unmarshal([]byte(s), &c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
	fmt.Println(c[1].First)
}

/* EX 3
- Partindo do código abaixo, utilize NewEncoder() e Encode() para enviar as informações como JSON para Stdout.
    - https://play.golang.org/p/BVRZTdlUZ_​
- Desafio: descubra o que é, e utilize, method chaining para conectar os dois métodos acima. */
func ex3() {
	u1 := character{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := character{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := character{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []character{u1, u2, u3}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(users)
	fmt.Println()
	encoder.Encode(user{"Guilherme", 28})
}

/* EX 4
- Partindo do código abaixo, ordene a []int e a []string.
	- https://play.golang.org/p/H_q75mpmHW */
func ex4() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

/* EX 5
- Partindo do código abaixo, ordene os []user por idade e sobrenome.
    - https://play.golang.org/p/BVRZTdlUZ_​
- Os valores no campo Sayings devem ser ordenados tambem, e demonstrados de maneira esteticamente harmoniosa. */
func printCharacters(cs []character) {
	for _, c := range cs {
		fmt.Println(c.First, c.Last, c.Age)
		for _, v := range c.Sayings {
			fmt.Println("-", v)
		}
	}
}

func ex5() {
	u1 := character{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := character{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := character{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	characters := []character{u1, u2, u3}

	// Sort sayings
	for _, c := range characters {
		sort.Strings(c.Sayings)
	}
	printCharacters(characters)

	sort.Sort(byAge(characters))
	printCharacters(characters)

	sort.Sort(byLast(characters))
	printCharacters(characters)

}

func main() {
	ex5()
}
