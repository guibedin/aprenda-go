package main

import "fmt"

/* EX 1
 Usando uma literal composta:
     - Crie um array que suporte 5 valores to tipo int
     - Atribua valores aos seus índices
- Utilize range e demonstre os valores do array.
- Utilizando format printing, demonstre o tipo do array. */
func ex1() {
	var array = [5]int{1, 2, 3, 4, 5}

	for _, v := range array {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", array)
}

/* EX 2
 Usando uma literal composta:
    - Crie uma slice de tipo int
    - Atribua 10 valores a ela
- Utilize range para demonstrar todos estes valores.
- E utilize format printing para demonstrar seu tipo. */
func ex2() {
	var s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", s)
}

/* EX 3
 Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
    - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
    - Do quinto ao último item do slice (incluindo o último item!)
    - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
    - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
	- Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item */
func ex3() {
	var s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
	fmt.Println(s[:3])
	// - Do quinto ao último item do slice (incluindo o último item!)
	fmt.Println(s[4:])
	// - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
	fmt.Println(s[1:7])
	// - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
	fmt.Println(s[2:9])
	// - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
	fmt.Println(s[2 : len(s)-1])
}

/* EX 4
- Começando com a seguinte slice:
    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Anexe a ela o valor 52;
- Anexe a ela os valores 53, 54 e 55 utilizando uma única declaração;
- Demonstre a slice;
- Anexe a ela a seguinte slice:
    - y := []int{56, 57, 58, 59, 60}
- Demonstre a slice x. */
func ex4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x, 52)
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

/* EX 5
- Comece com essa slice:
    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Utilizando slicing e append, crie uma slice y que contenha os valores:
	- [42, 43, 44, 48, 49, 50, 51] */
func ex5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := append(x[:3], x[6:]...)

	fmt.Println(y)
}

/* EX 6
- Crie uma slice usando make que possa conter todos os estados do Brasil.
	- Os estados: "Acre", "Alagoas", "Amapá", "Amazonas",
	"Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão",
	"Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará",
	"Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro",
	"Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima",
	"Santa Catarina", "São Paulo", "Sergipe", "Tocantins"
- Demonstre o len e cap da slice.
- Demonstre todos os valores da slice *sem utilizar range.* */
func ex6() {
	var estados = make([]string, 0, 26)
	estados = append(estados, []string{"Acre", "Alagoas", "Amapá", "Amazonas",
		"Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão",
		"Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará",
		"Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro",
		"Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima",
		"Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}...)

	fmt.Println(len(estados), cap(estados))

	for i := 0; i < len(estados); i++ {
		fmt.Printf("%v ", estados[i])
	}
	fmt.Println()
}

/* EX 7
- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados. */
func ex7() {
	var s = [][]string{
		[]string{
			"Guilherme",
			"Bedin",
			"PC",
		},
		[]string{
			"Camila",
			"Zago",
			"Dentista",
		},
	}

	fmt.Println(s)
}

/* EX 8, 9 e 10
- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.
- Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.
*/

func ex8() {
	m := map[string][]string{}

	m["Bedin_Guilherme"] = []string{"Hobby 1", "Hobby 2"}
	m["Zago_Camila"] = []string{"Hobby 3", "Hobby 4"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "Bedin_Guilherme")
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main() {

	ex8()
}
