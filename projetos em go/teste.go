package main

import "fmt"

func main() {

	ultimoteste(ultimoargumento)

}

func ultimoargumento() {
	fmt.Println("\nUltimo porque meus dedos estão doendo")
	fmt.Println("Não faço mais ideia do que digitar")
	fmt.Println("É isso acabou, até mais")
}
func ultimoteste(ultimo func()) {
	fmt.Println("É o ultimo por hora")
	fmt.Println("Ultimo até eu precisar usar isso de novo")
	fmt.Println("Ou até que tenha esquecido")
	ultimo()
}
