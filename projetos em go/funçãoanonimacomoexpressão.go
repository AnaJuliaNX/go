package main

import "fmt"

func main() {

	x := 10

	func(x int) {
		fmt.Println(x, "vezes 10 é:")
		fmt.Println(x * 10)
	}(x)

}

//Tela
//10 vezes 10 é:
//100


///////////////////////////////////////////////////////


package main

import "fmt"

func main() {

	x := 10
	y := func(x int) int {
		return x * 10
	}
	fmt.Println(x, "vezes 10 é:", y(x))
}
//tela
//10 vezes 10 é: 100
