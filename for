package main

import "fmt"

func main() {

	for mes := 1; mes <= 12; mes++ {
		for x := 1; x <= 30; x++ {
			fmt.Println("O mês é: ", mes, "e o dia é: ", x)
		}
	}

}
