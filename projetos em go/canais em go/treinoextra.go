//Teste para entender melhor como cada parte do processo funciona

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//aqui onde os canais vão convergir para um só
	canaldeteste := converge(trabalho("Isso é um treino"), trabalho("Isso também é"))
	//Aqui dou um limite de quantas vezes repetir, nesse caso, vai repetir 15 vezes, sendo que elas intercalam
	//uma pode repetir mais vezes que a outra e vice versa
	for repete := 0; repete < 10; repete++ {
		fmt.Println(<-canaldeteste)
	}
}
func trabalho(s string) chan string {
	canal := make(chan string)
	go func(s string, c chan string) {
		for i := 1; ; i++ {
			c <- fmt.Sprintf("Avisos paroquiais %v repetiu %v vezes", s, i)
			//isso aqui faz ter um intervalo de tempo entre cada execução de um segundo
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		}
	}(s, canal)
	return canal
}
func converge(x, y chan string) chan string {
	novo := make(chan string)
	go func() {
		for {
			novo <- <-x
		}
	}()
	go func() {
		for {
			novo <- <-y
		}
	}()
	return novo
}
