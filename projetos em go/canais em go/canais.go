//São o jeito certo de fazer sincronização e código concorrente
//Eles nos permitem transmitir valores entre goroutines
//Servem pra coordenar, sincronizar , orquestrar e buffering
//buffer: é como se eu deixasse um valor lá dentro que pode ser tirado depois
//EX: canal := make(chan int, 1), roda do mesmo jeito mas não é recomendável ficar usando ele. poder dar BO
//Se tiver mais de um canal e quiser usar o buffer preciso incluir todos os canais no buffer a menos que
//estejam dentro de uma func ai coloco só o que precisa

package main

import "fmt"

func main() {

	canal := make(chan int, 1)

	go func() {
		canal <- 42
		canal <- 43

	}()

	fmt.Println(<-canal)
	fmt.Println(<-canal)

}

///////////////////////////////////////////////////////////////////////////
//CANAIS BIDIRECIONAIS

//O QUE AS SETAS SIGNIFICAM:
//chan<-: coloca as informações no canal.
//<-chan: tira as informações do canal.

package main

import "fmt"

func main() {

	canal := make(chan int) //aqui ele ainda é um canal bidirecional

	go send(canal) //pelo menos um precisa ser uma goroutine

	receive(canal)
}

func send(s chan<- int) {
	s <- 14
//esse canal serve para colocar valores dentro do canal e coloquei o 14 no caso
}
func receive(r <-chan int) {
	fmt.Println("O valor recebido do canal foi: ", <-r)
//esse canal serve para receber/tirar informações
}

//isso tudo quer dizer que: no canal send eu coloquei o número 14 e no canal receive eu tirei a informação para
//exibir na tela, só reparar nas setas.

///////////////////////////////////////////////////////////////
//MAIS SOBRE CANAIS

package main

import "fmt"

func main() {

	c := make(chan int)
	cr := make(<-chan int) //receive
	cs := make(chan<- int) //send

	//ISSO AQUI NÃO PODERIA FAZER:
	// cr = cs
	//Um especifico não pode ser outro canal especifico

	fmt.Println("------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("c\t%T\n", cr) //isso exibe a tipagem dele
	fmt.Printf("c\t%T\n", cs)

	fmt.Println("------")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))

}

//ISSO TAMBÉM NÃO PODERIA FAZER:
//fmt.Printf("c\t%T\n", (chan int)(cr))
//fmt.Printf("c\t%T\n", (chan int)(cs))
//de especifico para geral não dá certo e nem pode
