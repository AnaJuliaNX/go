package main

import "fmt"

func main() {

	x := 2

	//Só exibe se forem as duas opções true
	if x%2 == 0 && x%3 == 0 {
		fmt.Println("Isso é multiplo de dois e três")
	}
	//Exibe se pelo menos uma das opções for true
	if x%2 == 0 || x%3 == 0 {
		fmt.Println("Isso é multiplo de dois ou três")
	}
}
