package main

import "fmt" 

func main() {
	x := 10
	if x > 100 {

		fmt.Println("Esse valor é maior que cem")

	} else if x < 10 {

		fmt.Println("Esse valor é menor que dez")

	} else {
		fmt.Println("Esse valor é dez")
	}
}
