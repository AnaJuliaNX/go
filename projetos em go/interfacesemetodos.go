package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}
type escritor struct {
	pessoa
	livros  int
	salário float64
}
type professor struct {
	pessoa
	anostrab int
	salário  float64
}

func (x escritor) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "já escrevi", x.livros, "livros", "e ouve só: Bom dia!")
}
func (x professor) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e ouve só: Bom dia!")
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()
}
func main() {
	srletra := escritor{
		pessoa: pessoa{
			nome:  "Rick Riordan",
			idade: 33,
		},
		livros:  32,
		salário: 10000022.1,
	}
	srprof := professor{
		pessoa: pessoa{
			nome:  "Juju",
			idade: 46,
		},
		anostrab: 18,
		salário:  4.511,
	}

	//fmt.Println(srletra)  //só exibe normal oq ta nos structs
	//fmt.Println(srprof)

	srletra.oibomdia()
	srprof.oibomdia() //cada um com seu método "oibomdia"

	serhumano(srletra)
	serhumano(srprof)     // tipod diferentes mas chamo a mesma coisa
}

//Tela
//Meu nome é Rick Riordan já escrevi 32 livros e ouve só: Bom dia!
//Meu nome é Juju e ouve só: Bom dia!
//Meu nome é Rick Riordan já escrevi 32 livros e ouve só: Bom dia!
//Meu nome é Juju e ouve só: Bom dia!
