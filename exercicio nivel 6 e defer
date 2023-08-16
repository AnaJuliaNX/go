package main

import "fmt"

func main() {
	fmt.Println(retornaint())
	fmt.Println(retorintstri())
}

func retornaint() int {
	return 10
}
func retorintstri() (int, string) {
	return 120, "livros lidos"
}

//tela
//10
//120 livros lidos

////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {

	si := []int{1, 2, 3, 4, 5, 6, 7, 8}
	is := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(função1(si...))
	fmt.Println(função2(is))
}

func função1(x ...int) int {
	total := 0
	for _, v := range x {    //parametro variadico do tipo int
		total += v
	}
	return total
}

func função2(x []int) int {
	total := 0
	for _, v := range x {
		total += v    //slice of int como argumento
	}
	return total
}
//tela
//36
//21

/////////////////////////////////////////////////////////

			EXEMPLO DE DEFER  
package main

import "fmt"

func main() {

	defer fmt.Println("Isso foi digitado antes")
	fmt.Println("Isso foi digitado depois")
}
//tela
//Isso foi digitado depois
//Isso foi digitado antes


////////////////////////////////////////////////////////

package main

import "fmt"

type livro struct {
	nome  string
	autor string
	nota  float64
}

func (mostra livro) exibe() {
	fmt.Println("O livro escolhido foi:", mostra.nome)
	fmt.Println("Do(a) autor(a):", mostra.autor)       //essa é a func (método) que eu criei do zero
	fmt.Println("E a nota que dei foi:", mostra.nota)
}
func main() {

	livro1 := livro{
		nome:  "Trono de vidro",
		autor: "Sarah J. Maas",
		nota:  9.9,
	}

	livro1.exibe() //mostra na tela os comando da func (método) que eu criei do zero
}

//tela
//O livro escolhido foi: Trono de vidro
//Do(a) autor(a): Sarah J. Maas
//E a nota que dei foi: 9.9

//////////////////////////////////////////////////////////////
        
			IMPLEMENTADO UMA INTERFACE
package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}
type circulo struct {
	raio float64
}

func (q quadrado) area() {
	total := q.lado * q.lado
	fmt.Println("A área do quadrado é: ", total)
}
func (c circulo) area() {
	total := math.Pi * 2 * c.raio
	fmt.Println("A área do circulo é: ", total)
}

type info interface {
	area()
}

func medida(i info) {
	i.area()
}
func main() {

	x := quadrado{
		lado: 4.0,
	}

	y := circulo{
		raio: 1.5,
	}

	medida(x)
	medida(y)
}
//tela
//A área do quadrado é:  16
//A área do circulo é:  9.42477796076938


/////////////////////////////////////////////////////////
		
			FUNÇÃO ANONIMA
package main

import "fmt"

func main() {

	func() {
		fmt.Println("Isso é uma função anonima")
	}()
}

//tela
//Isso é uma função anonima


////////////////////////////////////////////////////////

//atribuindo função á uma variável 

package main

import "fmt"

func main() {

	x := func() {
		fmt.Println("Isso é uma função anonima")
	}
	x()
}


//////////////////////////////////////////////////////////

		   FUNÇÃO QUE RETORNA FUNÇÃO
package main

import "fmt"

func main() {

	nomealeatorio := testedefunc()
	nomealeatorio()  //isso serve para chamar a função e executar tudo que foi criado nela
}
func testedefunc() func() {  //função que retorna uma função
	return func() {
		fmt.Println("Isso é uma função que retorna uma função")
		fmt.Println(true)
		fmt.Println(14)        //isso é o qur foi criado nela
		fmt.Println(14.5)
	}
}
//tela
//Isso é uma função que retorna uma função
//true
//14
//14.5

//////////////////////////////////////////////////////////

	FUNÇÃO COMO ARGUMENTO PARA OUTRA FUNÇÃO

package main

import "fmt"

func main() {

	recebeargumento(oargumento)
}

func oargumento() {      //passo essa como argumento para a de baixo
	fmt.Println("Esse é o argumento")
}

func recebeargumento(x func()) { //tem como parametro receber uma função

	fmt.Println("Esse recebe o argumento")

	x() //chamou a função do parametro
}
//tela
//Esse recebe o argumento
//Esse é o argumento
