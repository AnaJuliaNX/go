// O oposto da convergencia, quando vou de um ou poucos canais para muitos
//de um modo simples, tenho uma coisa e abro ela em milhares de outras coisas, no final elas se juntam em um 
//unico lugar (convergencia) para serem exibidas na tela

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	canal1 := make(chan int)
	canal2 := make(chan int)

	go manda(20, canal1)
	go outra(canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}
}

//mandamos x numeros para o primeiro canal
func manda(n int, canal chan int) {
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
}

//nessa segunda função ela pega cada item que foi no primeiro canal e manda uma goroutine pra cada um
func outra(canal1, canal2 chan int) {
	var wg sync.WaitGroup

//ou seja, eu tenho 20 itens no canal1 e para cada um deles eu vou chamar uma goroutine
//quando todas acabam de ser chamadas elas vão para o canal
	for v := range canal1 {
		wg.Add(1)
		go func(x int) {
			canal2 <- trabalho(x)
			wg.Done()
		}(v)
	}
//nesse momento eu tenho 20 goroutines trabalhando
	wg.Wait()
	close(canal2)
}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	return n
}

///////////////////////////////////////////////////////////////
//EXEMPLO DOIS

// nesse caso aqui eu tenho cinco go routines executando por vez
// então elas executando o que precisa até acabar, depois ais cinco fazem o mesmo e assim por diante
//ressaltando que ele executa de forma um pouco ou muito aleatória

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)
	funçao := 5 //as funções vão de cinco em cinco

	go manda(100, canal1) //são cem funções ao todo
	go outra(funçao, canal1, canal2)
	for v := range canal2 {
		fmt.Println(v)
	}
}

func manda(n int, canal chan int) {
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
}

func outra(funçao int, canal1, canal2 chan int) {
	var wg sync.WaitGroup
	for i := 0; i < funçao; i++ {
		wg.Add(1)
		go func() {
			for v := range canal1 {
				canal2 <- trabalho(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(canal2)
}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * 1000) //tem um intervalo de um segundo entre cada execução
	return n
}

//linha do tempo de como funciona: tinha uma coisa -> virou 5 goroutines -> trabalham o que precisa -> encerram
//Essas cinco coisas vão trabalhando até serem colocadas em uma unica coisa e exibida na tela

