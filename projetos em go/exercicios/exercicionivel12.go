// Package idade é apenas um exercicio feito em um curso de go lang. Fiz algumas alterações e simplifiquei ele
package idade

import "fmt"

// var anos int vai salvar o valor numérico que o usuario vai digitar
var anos int

func idade() {

	//fmt.Println nesse caso é para pedir a iformação ao usuário
	fmt.Println("Digite a idade em anos humanos do cachorro:")

	//fmt.Scan serve para armazenar o conteúdo digitado
	fmt.Scan(&anos)

	//total dá o resultado do calculo entre o valor digitado multiplicado por setee
	total := anos * 7

	//fmt.Println nesse caso exibe o resultado do calculo  e uma mensagem informativa
	fmt.Println("A idade em anos caninos do cachorro é: ", total)
}
