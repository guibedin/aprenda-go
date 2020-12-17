package main

import (
	"fmt"
)

/* EX 1
- Crie um valor e atribua-o a uma variável.
- Demonstre o endereço deste valor na memória. */
func ex1() {
	x := 10
	fmt.Println(x, &x)
}

/* EX 2
- Crie um struct "pessoa"
- Crie uma função chamada mudeMe que tenha *pessoa como parâmetro. Essa função deve mudar um valor armazenado no endereço *pessoa.
    - Dica: a maneira "correta" para fazer dereference de um valor em um struct seria (*valor).campo
    - Mas consta uma exceção na documentação. Link: https://golang.org/ref/spec#Selectors​
    - "As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method),
    - → x.f is shorthand for (*x).f." ←
    - Ou seja, podemos usar tanto o atalho p1.nome quanto o tradicional (*p1).nome
- Crie um valor do tipo pessoa;
- Use a função mudeMe e demonstre o resultado. */
type pessoa struct {
	nome  string
	idade int
}

func mudeMe(p *pessoa) {
	(*p).nome = "Camila"
	(*p).idade = 27
}

func ex2() {
	p := pessoa{"Guilherme", 28}
	fmt.Println(p)
	mudeMe(&p)
	fmt.Println(p)
}

func main() {
	ex2()
}
