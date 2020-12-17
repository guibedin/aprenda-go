/* EX 1
- Põe na tela: todos os números de 1 a 10000. */

/* EX 2
- Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
- Por exemplo:
    65
        U+0041 'A'
        U+0041 'A'
        U+0041 'A'
    66
        U+0042 'B'
        U+0042 'B'
        U+0042 'B'
	...e por aí vai. */

/* EX 3
- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu. */

/* EX 4
- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu. */

/* EX 5
- Demonstre o resto da divisão por 4 de todos os números entre 10 e 100 */

/* EX 8
- Crie um programa que utilize a declaração switch, sem switch statement especificado. */

/* EX 9
- Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito". */

package main

import (
	"fmt"
)

func ex1() {
	for i := 1; i <= 10000; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")
}

func ex2() {
	for i := 65; i <= 90; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U ", i)
		}
		fmt.Printf("\n")
	}
}

func ex3() {
	ano := 1992
	for ano < 2021 {
		fmt.Printf("%d ", ano)
		ano++
	}
	fmt.Printf("\n")
}

func ex4() {
	ano := 1992
	for {
		if ano == 2021 {
			break
		}
		fmt.Printf("%d ", ano)
		ano++
	}
	fmt.Printf("\n")
}

func ex5() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
	fmt.Printf("\n")
}

func ex8() {

	x := 6

	switch {
	case x == 5:
		fmt.Println("x 5")
	case x == 6:
		fmt.Println("x 6")
	case x == 7:
		fmt.Println("x 7")
	}
}

func ex9() {
	esporteFavorito := "basquete"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("futebol")
	case "basquete":
		fmt.Println("basquete")
	}
}

func main() {
	ex9()
}
