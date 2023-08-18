package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup  //PRECISO DISSO PARA CONSEGUIR EXECUTAR O WAITGROUP

func main() {

	fmt.Println(runtime.NumCPU())  //ISSO ME PERMITE VER QUANTOS PROCESSADORES TEM
	fmt.Println(runtime.NumGoroutine()) //QUANTAS GOROUTINES TEM INICIALMENTE

	wg.Add(2) //DIGO PARA OS WAITGROUPS QUANTAS GOROUTINES TENHO PARA EXECUITAR

	go func1()  //COLOCO O GO PARA QUE ISSO VIRE UMA GOROUTINE
	go func2()

	fmt.Println(runtime.NumGoroutine()) //QUANTAS GOROUTINES TENHO DEPOIS DE COLOCAR MAIS DUAS

	wg.Wait()  //DIGO PARA O PROGRAMA NÃO ENCERRAR AINDA PORTE TENHO GOROUTINE PARA EXECUTAR

}

func func1() {
	for i := 0; i < 5; i++ {
		fmt.Println("func1:", i)
		time.Sleep(20)  //ESNTRE CADA FUNÇÃO ELA DORME POR 20 MILISEGUNDOS ANTES DE EXECUTAR A PRÓXIMA
	}
	wg.Done() //DIGO PARA O PROGRAMA QUE A FUNÇÃO JÁ FOI EXECUTADA
}

func func2() {
	for i := 0; i < 5; i++ {
		fmt.Println("func2:", i)
		time.Sleep(20)
	}
	wg.Done()  //DIGO PARA O PROGRAMA QUE A FUNÇÃO JÁ FOI EXECUTADA
}


//Tela

//8  //NUMERO DE PROCESSSDORES
//1  //GOROUTINES INICIALMENTE, NO CASO SÓ A FUNÇÃO MAIN
//3  //GOROUTINES DPS DE EXECUTAR O PROGRAMA, NO CASO A MAIN, FUNC1 E FUNC2

//func2: 0
//func1: 0
//func1: 1
//func2: 1      //TA EXECUTANDO DE FORMA INTERCALADA, PARALELA, OU SEJA, CONCORRENTE
//func2: 2
//func1: 2
//func2: 3
//func1: 3
//func1: 4
//func2: 4
