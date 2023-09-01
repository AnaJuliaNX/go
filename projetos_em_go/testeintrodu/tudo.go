/*package main

import (
	"fmt"
)

func main() {
	//A melhor dica para fazer os teste é colocar as funções e testes na mesma ordem para não se perder
	a := Soma(1, 2, 3)
	b := Multiplica(2, 2)
	c := Numpi(3.14)
	d := Nome("Pelo jeito pode string")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

// Soma, essa função faz a soma entre três números previamente escolhidos
func Soma(i ...int) int {
	total := 0
	for _, v := range i {
		total = total + v
	}
	return total
}

// Multiplica, essa função faz a multiplicação de dois números previamente escolhidos
func Multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}

// Numpi mostra na tela o valor de Pi
func Numpi(i ...float64) float64 {
	total := 3.14
	for _, v := range i {
		total = v
	}
	return total
}

// Nome mostra na tela o conteúdo da string
func Nome(i ...string) string {
	fala := "Pelo jeito pode string"
	for _, v := range i {
		fala = v
	}
	return fala
}
*/

////////////////////////////////////////////////////////////////////////
//TESTE EM TABELA

package main

import "fmt"

func main() {

	a := Soma(2, 3, 5)
	b := Nome("Aelin")
	c := Simplao("Teste simples")
	d := Multiplica(10, 10)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func Soma(i ...int) int {
	total := 0
	for _, v := range i {
		total = total + v
	}
	return total
}

func Nome(i ...string) string {
	nome := "Aelin"
	for _, v := range i {
		nome = v
	}
	return nome
}

func Simplao(i ...string) string {
	simples := "Teste apenas"
	for _, v := range i {
		simples = v
	}
	return simples
}

func Multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}
