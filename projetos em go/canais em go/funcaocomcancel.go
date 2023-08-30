package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("erro checado 1:", ctx.Err())
	fmt.Println("numero de goroutines 1:", runtime.NumGoroutine()) //aqui vai ser uma pq só conta a func main

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done(): //se eu rodar o cancel e ele der um done, acaba aqui
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("funcionando", n) //se não, ele vai continuar exibindo na tela a mensagem "funinanndo"
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("erro checado 2:", ctx.Err())
	fmt.Println("numero de goroutines 2:", runtime.NumGoroutine()) //aqui é duas porque conta a go fun e a func main

	fmt.Println("prestes a cancelar o contexto") //avisa isso um pouco antes de fazer o cancelamento
	cancel()
	fmt.Println("contexto cancelado") //aqui cancela o contexto
//depois disso não roda mais nenhuma goroutine

	time.Sleep(time.Second * 2)
	fmt.Println("erro checado 3:", ctx.Err())
	fmt.Println("numero de goroutines 3:", runtime.NumGoroutine()) //volta a ser uma porque cancelei a go func
}

////////////////////////////////////////////////////////////////////////
//EXEMPLO 2

//ficamos esperando uma certa condição, supriu a condição cancela e acabou

package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //cancela quando estiver finalizado

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 { //ela vai mostrar esse code block, entaõ, quando terminar de exibir os cinco numeros ela cancela
			break
		}
	}
}
func gen(ctx context.Context) <-chan int {
	dts := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():  //tem uma go func esperando um sinal de cancel
				return //quando recebe o cancel ela morre
			case dts <- n:
				n++  //se não morrer ela incrementa um contador
			}
		}
	}()
	return dts
}

////////////////////////////////////////////////////
//EXEMPPLO TRÊS, COM DEADLINE

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
"time.Now" pega do momento atual em que vai rodar e add 50 milissegundos
	d := time.Now().Add(50 * time.Millisecond)  //se mudasse para "second" ele rodava
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()
//vou rodar o code block de baixo e depois cancelar tudo

	select {
	case <-time.After(1 * time.Second):
//isso nunca vai executar porque a minha função aqui ta pra 1 milissegundo
//ou seja, o que to esperando jamais vai rodar e vou estourar minha deadline
		fmt.Println("dormiu demais")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

///////////////////////////////////////////////////////////////
//EXEMPLO QUATRO, COM WITHTIMEOUT

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
//vai acontecer o mesmo com o deadline, vou estourar ele
//diferente do doutro que especifico o tempo, nesse eu só desconto 50 milisegundos
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)//se mudasase pra second ele rodava
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("dormiu demais")
	case <-ctx.Done():
			fmt.Println(ctx.Err())
	}
}

////////////////////////////////////////////////////
//EXEMPLO CINCO, COM WITHVALUE

package main

import (
	"context"
	"fmt"
)

func main() {
	//como o nome já diz, ele me pasa um valor     isso aqui é como um valor referente ao outro
	//no caso eu vou pedir um e ele me da o da frente
	ctx := context.WithValue(context.Background(), "Linguagem", "Go")
	foo(ctx, "Linguagem")

	ctx = context.WithValue(ctx, "Gato de", "Botas")
	foo(ctx, "Gato de")

	//se eu quisesse que a key "cor" tivesse algum valor eu teria que colocar o comandoi de baixo
	//ctx = context.WithValue(context.Background(), "cor", "roxo")
	foo(ctx, "cor") //vou pedir uma cor, mas como não tenho nenhuma ele não encontra nenhuma
}

func foo(ctx context.Context, s string) {
	if v := ctx.Value(s); v != nil {
		fmt.Println("Valor perdido:", v)  //isso aqui é o que vou exibir na tela caso tenha o valor
		return
	}
	fmt.Println("key não encontrada:", s) //isso é o que vou exibir caso não tenha o valor
}
