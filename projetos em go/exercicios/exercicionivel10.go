//EXERCICIO UM
//fazer o código funcionar usando uma função anonima auto-executavel e depois usando buffer

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

/////////////////////////////////////////////////
//EXERCICIO UM, BUFFER

package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1) //só adicionei o valor 1

	c <- 42

	fmt.Println(<-c)
}

/////////////////////////////////////////////////
//EXERCICIO DOIS

package main

import (
	"fmt"
)

func main() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("---------\n")
	fmt.Printf("cs\t%T\n", cs)
}

////////////////////////////////////////////////
//EXERCICIO TRÊS

package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("quase saindo")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(r <-chan int) {
	for v := range r {
		fmt.Println(v)
	}
}

///////////////////////////////////////////////////
//EXERCICIO QUATRO

package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("quase saindo")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		q <- 42
	}()
	return c
}

func receive(c <-chan int, q chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("O conteudo do canal é:", v)
		case <-q:
			return
		}
	}
}

///////////////////////////////////////////////////
//EXERCICIO CINCO

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 15
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

///////////////////////////////////////////////////
//EXERCICIO SEIS

package main

import (
	"fmt"
)

func main() {
	canalaqui := make(chan int)
	go canal(canalaqui)

	for v := range canalaqui {
		fmt.Println("Acho que deu certo:", v)
	}
}
func canal(canalaqui chan int) {
	for i := 1; i <= 100; i++ {
		canalaqui <- i
	}
	close(canalaqui)
}

///////////////////////////////////////////////////
//EXERCICIO SETE

package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)
	for i := 0; i < 10; i++ {

		go func() {
			for j := 0; j < 10; j++ {
				canal <- j
			}
		}()
	}
	for k := 0; k < 100; k++ {
		fmt.Println(k, "\t", <-canal)
	}
}
