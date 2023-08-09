package main

import "fmt"

func main() {

	amigos := "teste"
	switch amigos {
	case "Ada":
		fmt.Println("A Ada sempre vem com o Ben")
		fallthrough
	case "Ben":
		fmt.Println("O Ben sempre vem com a Ada")
	case "Evie":
		fmt.Println("A Evie sempre vem com o Leo")
		fallthrough
	case "Leo":
		fmt.Println("O Leo sempre vem com a Evie")
	default:
		fmt.Println("Ninguém veio ainda ou já foram embora")
	}

}

//O fallthrough ignora o próximo case e exibe o próximo print
