//É como o switch pra canais e não é sequencial
//ele bloqueia até ter um valor que bate com os casos especificados, quando recebe ele executa
//se tiver mais de um caso válido ele escolhe um aleatório para executar, não tenho controle sobre isso

//package main

//import "fmt"

//func main() {

//	a := make(chan int)
//vários canais
//	b := make(chan int)
//	x := 500

//	go func(x int) {
//		for i := 0; i < x; i++ {
//			a <- i
//		}
//	}(x / 2)

//	go func(x int) {
//		for i := 0; i < x; i++ {
//			b <- i
//		}
//	}(x / 2)

//	for i := 0; i < x; i++ {
//		select {  //eu posso usar esse select para receber informações de vários canais em um mesmo lugar
//		case v := <-a:
//			fmt.Println("Canal A: ", v)
//aqui no caso é o meu "switch com cases"
//		case v := <-b:
//			fmt.Println("Canal B: ", v)
//		}
//	}
//}

////////////////////////////////////////////////////////////////////////////////
//EXEMPLO DOIS DE SELECT

//package main

//import "fmt"

//func main() {

//	canal := make(chan int)
//	quit := make(chan int)
//	go recebeQuit(canal, quit)
//	enviaPCanal(canal, quit)

//}

//func recebeQuit(canal chan int, quit chan int) {
//	for i := 0; i < 50; i++ {
//		fmt.Println("Recebido: ", <-canal)
//	}
//	quit <- 0
//}

//func enviaPCanal(canal chan int, quit chan int) {
//	algumacoisa := 1
//	for {
//		select {
//		case canal <- algumacoisa:  //manda coisa para um canal
//			algumacoisa++

//		case <-quit: //recebe coisa de um canal
//			return
//		}

//	}
//}

////////////////////////////////////////////////////////////////////////////////
//EXEMPLO TRES DE SELECT

package main

import "fmt"

func main() {

	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	go mandaNumero(par, impar, quit)

	receive(par, impar, quit)
}

func mandaNumero(par, impar, chan int, quit, chan bool) {
	total := 10
	for i := 0; i < total; i++ {
		if i % 2 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
	quit <- true
}

func receive(par, impar, chan int, quit, chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("O numero: ", v, "é par")

		case v := <-impar:
			fmt.Println("O numero: ", v, "é impar")

		case <-quit:

		}
	}

}
