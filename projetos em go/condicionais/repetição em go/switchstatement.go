package main

import "fmt"

func main() {
	esportefavorito := "carro"

	switch esportefavorito {
	case "volei":
		fmt.Printf("Você é perfeito(a)")
	case "futebol":
		fmt.Printf("É um esporte padrão")
	case "basquete":
		fmt.Printf("Você é meio esquisito")
	case "handebol":
		fmt.Printf("Nada a declarar")
	case "tenis":
		fmt.Printf("Esporte de classe média alta")
	case "golf":
		fmt.Printf("Você é um idoso tenho certeza")
	default:
		fmt.Printf("Esporte não listado, tente outro")
	}
}
