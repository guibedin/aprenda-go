/* EX 1
- Desafio surpresa!
- Format printing:
    - Decimal       %d
    - Hexadecimal   %#x
    - Unicode       %#U
    - Tab           \t
    - Linha nova    \n
- Faça um loop dos números 33 a 122, e utilize format printing para demonstrá-los como texto/string. */

package main

import (
	"fmt"
)

func ex1() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%c ", i)
	}
	fmt.Printf("\n")
}

func main() {

	ex1()
}
