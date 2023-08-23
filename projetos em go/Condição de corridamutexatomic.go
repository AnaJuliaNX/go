                        CONDIÇÃO DE CORRIDA

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())  //ISSO FAZ COM QUE EU VEJA QUANTAS CPUS ESTOU USANDO
	fmt.Println("Goroutines:", runtime.NumGoroutine())  //ISSO FAZ COM QUE EU VEJA QUANTAS GO ROUTINES TENHO

	contador := 0
	totaldegoroutines := 10

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			v := contador
			runtime.Gosched() //FAZ PARAR UMA GO ROUTINE E IR PARA A OUTRA, PARECIDO COM UM TIME SLEEP
			v++
			contador = v    //ISSO SÓ FAZ SOMAR MAIS UM NÚMERO NOS DIGITOS DA GO ROUTINE
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)
}

//tela
//CPUs: 8
//Goroutines: 1
//Goroutines: 2
//Goroutines: 2
//Goroutines: 2
//Goroutines: 2
//Goroutines: 3
//Goroutines: 4
//Goroutines: 5
//Goroutines: 6
//Goroutines: 2
//Goroutines: 2
//Goroutines: 1
//Valor final: 9  //ESSE VALOR É ALTERADO TODA VEZ QUE RODO O CÓDIGO NOVAMENTE
  //É UM VALOR ALEATÓRIO


///////////////////////////////////////////////////////
                          MUTEX


    package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	contador := 0
	totaldegoroutines := 10

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	var mu sync.Mutex  //USO ISSO PRA EXECUTAR APENAS UMA COISA POR VEZ 

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			mu.Lock() //COMO SE BLOQUEASSE PARA COMEÇAR A EXECECUTAR UM POR VEZ 
			v := contador
			runtime.Gosched()
			v++
			contador = v 
			mu.Unlock()  //COMO SE DESBLOQUEASSE
			wg.Done() 
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)
}

//TELA
//CPUs: 8
//Goroutines: 1
//Goroutines: 2
//Goroutines: 2
//Goroutines: 3
//Goroutines: 2
//Goroutines: 3
//Goroutines: 2
//Goroutines: 2
//Goroutines: 3
//Goroutines: 4
//Goroutines: 5
//Goroutines: 1
//Valor final: 10  //VAI EXIBIR CERTO O TANTO DE GOROUTINES QUE TENHO




///////////////////////////////////////////////////////
                          ATOMIC

  package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic" //PRECISA DESSE PACOTE PAR DAR CERTO
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())

	var contador int64  //PRECISO CRIAR UMA VARIAVEL 
	contador = 0
	totaldegoroutines := 10

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1) //ELE USA PONTEIROS PARA EXECUTAR 
			runtime.Gosched()
			fmt.Println("Contador:\t", atomic.LoadInt64(&contador))
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Valor final:", contador)
}


//TELA 

//CPUs: 8
//Goroutines: 2
//Goroutines: 3
//Contador:	 1
//Goroutines: 4
//Goroutines: 4
//Goroutines: 5
//Goroutines: 6
//Goroutines: 7
//Contador:	 7
//Contador:	 2
//Contador:	 3
//Contador:	 4
//Contador:	 6
//Contador:	 5
//Contador:	 8
//Goroutines: 8
//Goroutines: 2
//Goroutines: 3
//Contador:	 9
//Contador:	 10
//Valor final: 10
