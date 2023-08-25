//Consiste, basicamente, em pegar dados de muitos pontos e convergir para poucos  ou um canal
//Foi o que fizemos no exemplo abaixo. Convergimos de dois canais (par e impart) para m só (converge)

//package main

//import (
//	"fmt"
//	"sync"
//)

//func main() {

//	par := make(chan int)
//	impar := make(chan int)
//	converge := make(chan int)

//	go envia(par, impar)
//	go recebe(par, impar, converge)

//aqui implementou as funcionalidades e exibiu tudo na tela
//	for v := range converge {
//		fmt.Println("Valor recebido: ", v)
//	}
//}

// aqui onde tem os dois canais (PAR E IMPAR) antes da convergenecia
//func envia(p, i chan int) {
//	x := 10
//	for n := 0; n < x; n++ {
//		if n%2 == 0 {
//			p <- n
//		} else {
//			i <- n
//		}
//	}
//	close(p)
//	close(i)
//}

// aqui é onde faz a convergencia dos dois canais para um só (CONVERGE)
//func recebe(p, i, c chan int) {
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go func() {
//		for v := range p {
//			c <- v
//		}
//		wg.Done()
//	}()
//	wg.Add(1)
//	go func() {
//		for v := range i {
//			c <- v
//		}
//		wg.Done()
//	}()
//	wg.Wait()
//	close(c)
//}

// ///////////////////////////////////////////////////////////
// SEGUNDO EXEMPLO
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//o converge vai pegar dois canais e criar outro com as duas informações juntas
	canal := converge(trabalho("Quero livro novo"), trabalho("Quero ir na sebo"))

	//aqui eu dou um limite de vezes que vai tirar a informação do canal
	//se não repete pra sempre ou até dar pau
	for x := 0; x < 16; x++ {
		fmt.Println(<-canal)
	}
}

func trabalho(s string) chan string {
	canal := make(chan string)
	go func(s string, c chan string) {
		for i := 1; ; i++ {
			//isso vaiu exibir meu recado e o valor de vezes que foi repetido
			c <- fmt.Sprintf("Recado importante: %v quero esse tanto %v", s, i)
			//esse faz com que entre uma execução e outra tenha um intervalo de um segundo
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(s, canal)
	return canal
}
func converge(x, y chan string) chan string {
	novo := make(chan string)
	go func() {
		for {
			novo <- <-x //tudo que estiver no canal "x" vai para o canal convergente "novo"
		}
	}()
	go func() {
		for {
			novo <- <-y //tudo que estiver no canal "y" vai para o canal convergente "novo"
		}
	}()
	return novo
}
