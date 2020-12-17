package main

import (
	"fmt"
)

/* EX 1
- Exercício:
    - Crie uma função que retorne um int
    - Crie outra função que retorne um int e uma string
    - Chame as duas funções
	- Demonstre seus resultados */
func ex1F1() int {
	return 5
}
func ex1F2() (int, string) {
	return 10, "f2"
}

func ex1() {

	x, y := ex1F2()
	fmt.Println(ex1F1(), x, y)
}

/* EX 2
- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função. */
func ex2F1(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

func ex2F2(x []int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

func ex2() {
	fmt.Println(ex2F1([]int{1, 2, 3, 4, 5}...))
	fmt.Println(ex2F2([]int{1, 2, 3, 4, 5}))
}

/* EX 3
- Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence. */
func ex3() {
	defer fmt.Println("Primeira linha da funcao (defer)")
	fmt.Println("Segunda linha da funcao")
}

/* EX 4
- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor. */
type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) mostra() {
	fmt.Println(p.nome, p.sobrenome, p.idade)
}

func ex4() {
	p1 := pessoa{"Guilherme", "Bedin", 28}
	p1.mostra()
}

/* EX 6
- Crie e utilize uma função anônima. */
func ex6() {
	func(x int) {
		fmt.Println(x)
	}(20)
}

/* EX 7
- Atribua uma função a uma variável.
- Chame essa função. */
func ex7() {
	f := func(x int) int {
		return x + 1
	}

	fmt.Println(f(10))
}

/* EX 8
- Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada. */
func ex8F1() func() {
	return func() {
		fmt.Println("funcao retornada")
	}
}

func ex8() {
	fr := ex8F1()
	fr()
}

/* EX 9
- Callback: passe uma função como argumento a outra função. */
func ex9F1(f func(x int), x int) {
	fmt.Println(x)
	f(x + 1)
}

func ex9() {
	ex9F1(func(x int) {
		fmt.Println("X no callback", x)
	}, 10)
}

/* EX 10
- Demonstre o funcionamento de um closure.
- Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope. */
func ex10F1() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func ex10() {
	f1 := ex10F1()
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())

	f2 := ex10F1()
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
}

func main() {
	ex10()
}
