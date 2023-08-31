package idade

import "fmt"

var anos int

func idade() {

	fmt.Println("Digite a idade em anos humanos do cachorro:")
	fmt.Scan(&anos)
	total := anos * 7
	fmt.Println("A idade em anos caninos do cachorro é: ", total)
}
