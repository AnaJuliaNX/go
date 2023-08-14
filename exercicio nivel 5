package main

import "fmt"

type veiculo struct {
	nportas int
	cor     string
}
type caminhonete struct {
	veiculo
	tração bool
}
type sedan struct {
	veiculo
	luxo bool
}

func main() {
	carro1 := caminhonete{
		veiculo: veiculo{
			nportas: 2,
			cor:     "vermelho",
		},
		tração: true,
	}
	carro2 := sedan{
		veiculo: veiculo{
			nportas: 4,
			cor:     "azul escuro",
		},
		luxo: false,
	}
	fmt.Println("Dados da caminhonete:")
	fmt.Println("portas, cor, tração")
	fmt.Println(carro1)

	fmt.Println("\nDados do sedan:")
	fmt.Println(" portas, cor, luxo")
	fmt.Println(carro2)

	fmt.Println("\nportas")
	fmt.Println(carro1.nportas) //serve para exibir apenas uma das opções dos structs
	fmt.Println(carro2.nportas)
}
//tela
//Dados da caminhonete:
//portas, cor, tração
//{{2 vermelho} true}

//Dados do sedan:
// portas, cor, luxo
//{{4 azul escuro} false}

//portas
//2
//4


//////////////////////////////////////////////////////////////////


package main

import "fmt"

func main() {
	livro := struct {
		nome        string
		autor       string
		páginas     []int
		notaemotivo map[float64]string
	}{

		nome:    "Jantar secreto",
		autor:   "Raphael Montes",
		páginas: []int{368},
		notaemotivo: map[float64]string{
			9.8: "Só não é 10 porque não foi tão uau assim",
		},
	}
	fmt.Println(livro)
}
//Tela
//{Jantar secreto Raphael Montes [368] map[9.8:Só não é 10 porque não foi tão uau assim]}
