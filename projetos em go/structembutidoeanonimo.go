package main

import "fmt"

type livro struct {
	nome     string
	genero   string
	adaptado bool
	nota     float64
}

func main() {

	livro1 := livro{
		nome:     "Jantar secreto,",
		genero:   "thriller,",
		adaptado: false,
		nota:     8.7,
	}

	livro2 := livro{
		nome:     "Jogos vorazes,", //duas formas de fazer
		genero:   "distopia,",
		adaptado: true,
		nota:     9.9,
	}

	livro3 := livro{"Trono de vidro,", "fantasia,", false, 10}
	fmt.Println(livro1)
	fmt.Println(livro2)
	fmt.Println(livro3)
}

//tela
//{Jantar secreto, Thriller, false 8.7}
//{Jogos vorazes, distopia, true 9.9}
//{Trono de vidro, fantasia, false 10}


/////////////////////////////////////////////////////

package main

import "fmt"

type livro struct {
	nome   string
	genero string
}

type personagem struct {
	livro
	adaptado bool
}

func main() {

	livro1 := personagem{
		livro: livro{
			nome:   "Jogos vorazes,",
			genero: "distopia",
		},
		adaptado: true,
	}
	livro2 := personagem{
		livro: livro{
			nome:   "Trono de vidro,",
			genero: "fantasia",
		},
		adaptado: false,
	}
	livro3 := personagem{livro{"Jantar secreto", "suspense"}, false}  //pode ser feito assim também

	fmt.Println("Nome, genero e se foi adaptado:")
	fmt.Println(livro1)
	fmt.Println(livro2)
	fmt.Println(livro3)
}
//tela
//Nome, genero e se foi adaptado:
//{{Jogos vorazes, distopia} true}
//{{Trono de vidro, fantasia} false}
//{{Jantar secreto suspense} false}


///////////////////////////////////////////////////////////

package main

import "fmt"

func main() {
	livro := struct {
		nome     string
		adaptado bool
	}{
		nome:     "Jogos vorazes,",
		adaptado: true}

	fmt.Println("Nome e se foi adaptado")
	fmt.Println(livro)
}
//tela
//Nome e se foi adaptado
//{Jogos vorazes, true}
