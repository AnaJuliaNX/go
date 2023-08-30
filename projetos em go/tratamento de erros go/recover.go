package main

import (
	"fmt"
)

func main() {
	f()
	fmt.Println("retorno normal do f")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			//caso o recover seja diferente de nada exibe o que escrevi embaixo
			fmt.Println("Recovered no f", r)
		}
	}()
	fmt.Println("chamando g")
	g(0) //começa em 0 e cada vez que ela chama ela mesma soma mais um número
	fmt.Println("retorno normal do g")
}

func g(i int) {
	if i > 3 { //quando chegar no g4 faz um panico, vai fechar todas as executadas anteriormente e
		//como a função f ta ligada na função g o panic vai tentar encerrar ela, mas, todavia,
		//entretanto ela tem um recover que é como um alerta dizendo que á ta tudo bem e não precisa encerrar
		fmt.Println("Em panico!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer no g", i)
	fmt.Println("Printando em g", i)
	g(i + 1) //função recursiva porque chama ela mesma com mais um, no caso soma ela e mais um
}

//TELA

//chamando g
//Printando em g 0
//Printando em g 1  //aqui executou todos do g de boas
//Printando em g 2
//Printando em g 3
//Em panico!        //aqui deu o panic
//Defer no g 3
//Defer no g 2       //começou a fechar todos por ordem do que abriu por ultimo
//Defer no g 1
//Defer no g 0
//Recovered no f 4   //achou o recover no f, voltou a ser normla sem o panic
//retorno normal do f   //e retornou normal para a função main
