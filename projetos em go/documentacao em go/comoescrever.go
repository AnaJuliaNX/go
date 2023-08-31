//Deve ser bem escrita e correta, mas de facil alteração e de manter

//Escrever sempre a documentação junto com o código, abre uma exceção caso seja muita documentação, ai cria
//outro arquivo comum nome padraão "doc.go".

//Sempre começa a frase com o nome do elemento a ser documentado.

//Para documentar um tipo, uma variável, uma constante ou um pacote escreva o documentário imediatamente antes
//da declaração do elemento, sem espaços em branco entre eles.

// Package testando serve para aplicar os metodos ensinados. Não é realmente util para nada
package testanto

import "fmt"

// X serev para medir a utilidade
const X = 4

// TEste serve só de exemplo
func Teste() {
	fmt.Println("Essa função tem apenas fins educacionais")
}

/*
teste2 começa com letra minuscula, ou sej,a não é exportado se eu quiser
 como não vai ser exportado ele não sera visivel na documentação
*/
func teste2() {
	fmt.Println("qualquer coisa além disso é inutil")
}

// Teste3 tem letra maiuscula e é a ultima coisa
func Teste3() {
	fmt.Println("Não perca tempo achando que tem muita utilidade. Nivel de utilidade: ", X)
}
