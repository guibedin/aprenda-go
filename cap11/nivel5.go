package main

import (
	"fmt"
)

/* EX 1
- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete. */
func ex1() {
	type pessoa struct {
		nome           string
		sobrenome      string
		saboresSorvete []string
	}

	pessoa1 := pessoa{
		nome:           "Guilherme",
		sobrenome:      "Bedin",
		saboresSorvete: []string{"Sabor 1", "Sabor 2"},
	}

	pessoa2 := pessoa{
		nome:           "Camila",
		sobrenome:      "Zago",
		saboresSorvete: []string{"Sabor 3", "Sabor 4"},
	}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, v := range pessoa1.saboresSorvete {
		fmt.Println(v)
	}
	fmt.Println()

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, v := range pessoa2.saboresSorvete {
		fmt.Println(v)
	}
}

/* EX 2
- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior. */
func ex2() {

	type pessoa struct {
		nome           string
		sobrenome      string
		saboresSorvete []string
	}

	pessoas := make(map[string]pessoa)

	pessoas["Bedin"] = pessoa{
		nome:           "Guilherme",
		sobrenome:      "Bedin",
		saboresSorvete: []string{"Sabor 1", "Sabor 2"},
	}

	pessoas["Zago"] = pessoa{
		nome:           "Camila",
		sobrenome:      "Zago",
		saboresSorvete: []string{"Sabor 3", "Sabor 4"},
	}

	for _, v := range pessoas {
		fmt.Println(v.nome, v.sobrenome)
		for _, v := range v.saboresSorvete {
			fmt.Println(v)
		}
	}
}

/* EX 3
- Crie um novo tipo: veículo
    - O tipo subjacente deve ser struct
    - Deve conter os campos: portas, cor
- Crie dois novos tipos: caminhonete e sedan
    - Os tipos subjacentes devem ser struct
    - Ambos devem conter "veículo" como struct embutido
    - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
    - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
- Usando os structs veículo, caminhonete e sedan:
    - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
    - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
- Demonstre estes valores.
- Demonstre um único campo de cada um dos dois. */
func ex3() {
	type veiculo struct {
		portas int
		cor    string
	}

	type caminhonete struct {
		veiculo
		tracaoNasQuatro bool
	}

	type sedan struct {
		veiculo
		modeloLuxo bool
	}

	caminhonete1 := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor:    "vermelho",
		},
		tracaoNasQuatro: true,
	}

	sedan1 := sedan{
		veiculo: veiculo{
			portas: 4,
			cor:    "azul",
		},
		modeloLuxo: false,
	}

	fmt.Println(caminhonete1)
	fmt.Println(caminhonete1.cor)
	fmt.Println(sedan1)
	fmt.Println(sedan1.modeloLuxo)
}

/* EX 4
- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice. */
func ex4() {
	sa := struct {
		numeros map[int]string
		nomes   []string
	}{
		numeros: make(map[int]string),
	}

	sa.nomes = append(sa.nomes, "Guilherme")
	sa.nomes = append(sa.nomes, "Camila")
	sa.numeros[0] = "Zero"
	sa.numeros[1] = "Um"

	fmt.Println(sa)
}

func main() {
	ex4()
}
