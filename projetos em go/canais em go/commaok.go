//EXEMPLO SIMPLES E INICIAL DE COMMA OK

package main

import "fmt"

func main() {

//canal := make(chan int)
//go func() {
//	canal <- 10
//	close(canal)
//}()

canal2 := make(chan int)
go func() {
	canal <- 15    //Se tivesse feito desse jeito adicionando mais um canal ai o "v, ok" funcionaria normal
	close(canal2)
}()

v, ok := <-canal
fmt.Println(v, ok)

v, ok = <-canal //Esse aqui vai exibir "0 false" porque não tenho nenhum canal para exbiri um valor
fmt.Println(v, ok)
}

//////////////////////////////////////////////////////////////
//ARRUMANDO O CÓDIGO DO EXERCICIO ANTERIOR COM COMMA OK

package main

import "fmt"

func main() {

	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	go mandaNumero(par, impar, quit)

	receive(par, impar, quit)
}

func mandaNumero(par, impar chan int, quit chan bool) {
	total := 100
	for i := 0; i < total; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
	quit <- true
}

func receive(par, impar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("O número", v, "é par")

		case v := <-impar:
			fmt.Println("O número", v, "é par")

		case v, ok := <-quit:  //aqui tá o comma ok
			if !ok {
				fmt.Println("Não deu certo, olha só", v)
			} else {
				fmt.Println("Encerrando. Recebemos:", v)
			}
			return
		}
	}
}
