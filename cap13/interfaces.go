package main

import (
	"fmt"
	"math"
)

/* EX 5
- Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo" */

type quadrado struct {
	lado float64
}

// Método para o tipo quadrado
func (q quadrado) area() float64 {
	return q.lado * q.lado
}

type circulo struct {
	raio float64
}

// Método para o tipo círculo
func (c circulo) area() float64 {
	return 2 * math.Pi * c.raio
}

// Interface - Tudo que implementar o método 'area()' é uma figura
type figura interface {
	area() float64
}

// Função que recebe a interface
func info(f figura) float64 {
	return f.area()
}

func ex5() {
	q := quadrado{5}
	c := circulo{2}

	fmt.Println(info(q))
	fmt.Println(info(c))
}

/*
func main() {
	ex5()
}*/
