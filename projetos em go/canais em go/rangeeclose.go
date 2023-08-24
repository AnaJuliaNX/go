//CUIDADOS PARA SE TER COM O RANGEE O CLOSE:
//O range vai funcionar como um loop então se não for fechado ele vai repetir pra sempre e dar BO
//Por isso usamos o close, ele fecha o loop do range quando a minha execução chega ao fim

package main

import "fmt"

func main() {

	c := make(chan int)

	go meuloop(10, c)

	for v := range c {  //range vai funcionar como um loop

		fmt.Println("Recebido do canal: ", v)
	}
}

func meuloop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s) //usa o close pra dizer pro canal que ele já pode parar de esperar se não ele fica
//esperando pra sempre receber mais alguma coisa
}

//////////////////////////////////////////////////////
//MEXENDO MAIS UM POUCO NISSO

package main

import "fmt"

func main() {

	c := make(chan int)

	go meuloop(10, c)
	prints(c)
}

func meuloop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s)
}

func prints(r <-chan int) {
	for v := range r {
		fmt.Println("Recebido do canal: ", v)
	}

}
