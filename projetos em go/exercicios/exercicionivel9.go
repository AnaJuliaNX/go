package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	novagoroutine(100)

	wg.Wait()

}

func novagoroutine(i int) {

wg.Add(i)
for j := 0; j < i; j++ {
x := j
go func(i int) {
	fmt.Println("Eu sou a goroutine numero: ", i+1)
	wg.Done()

}(x)

}
}

/////////////////////////////////////////////////////////////////////////////
//EXERCICIO DOIS

package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) falar() {
	fmt.Println(p.nome, "Diz olá")
}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {

p1 := pessoa{
	nome:  "Aelin",
	idade: 21,
}

p1.falar() //isso aqui é um shortcut (recorte curto/atalho) pra isso aqui (&p1).falar()
(&p1).falar() //isso aqui é a maneira "mais correta" de se escrever

dizerAlgumaCoisa(&p1)
//dizerAlgumaCoisa(p1) assim são funciona porque não tenho um ponteiro para p1

//o com ponteiro funciona porque o ponteiro vai mostrar o type pessoa que é integrado na interface humanos
// que por sua vez faz parte da função dizerAlgumaCoisa
//}

/////////////////////////////////////////////////////////////////////////////
//EXERCICIO TRES E QUATRO

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs: ", runtime.NumCPU())

	contador := 0
	totaldegoroutine := 10

	var wg sync.WaitGroup
	wg.Add(totaldegoroutine)

	var mu sync.Mutex

	for i := 0; i < totaldegoroutine; i++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
		//fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Valor final: ", contador)
}
