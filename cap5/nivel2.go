/* EX 1
- Escreva um programa que mostre um número em decimal, binário, e hexadecimal. */

/* EX 2
- Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
	- ==
	- !=
	- <=
	- <
	- >=
	- >
- Demonstre estes valores. */

/* EX 3
- Crie constantes tipadas e não-tipadas.
- Demonstre seus valores. */

/* EX 4
- Crie um programa que:
    - Atribua um valor int a uma variável
    - Demonstre este valor em decimal, binário e hexadecimal
    - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
    - Demonstre esta outra variável em decimal, binário e hexadecimal */

/* EX 5
- Crie uma variável de tipo string utilizando uma raw string literal.
- Demonstre-a. */

/* EX 6
- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores. */

package main

import (
	"fmt"
)

// EX 3
const x int = 10
const y = 10

// EX 6
const (
	_  = iota
	a1 = iota + 2020
	a2
	a3
	a4
)

func ex1() {
	x := 27
	fmt.Printf("%d %b %#x\n", x, x, x)
}

func ex2() {
	a := (10 == 10)
	b := (10 != 10)
	c := (10 <= 10)
	d := (10 < 10)
	e := (10 >= 10)
	f := (10 > 10)

	fmt.Println(a, b, c, d, e, f)
}

func ex3() {
	fmt.Println(x, y)
}

func ex4() {
	x := 10
	fmt.Printf("%d %b %#x\n", x, x, x)

	y := x << 1
	fmt.Printf("%d %b %#x\n", y, y, y)
}

func ex5() {
	s := `LITERAL
	raw					string
		.`

	fmt.Println(s)
}

func ex6() {
	fmt.Println(a1, a2, a3, a4)
}

func main() {
	ex6()
}
