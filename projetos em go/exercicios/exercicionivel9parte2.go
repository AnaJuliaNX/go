//EXERCICIO CINCO

//package main

//import (
//	"fmt"
//	"runtime"
//	"sync"
//	"sync/atomic"
//)

//func main() {

//	fmt.Println("CPUs: ", runtime.NumCPU())

//	var contador int64
//	contador = 0
//	totaldegoroutine := 15

//	var wg sync.WaitGroup
//	wg.Add(totaldegoroutine)

//	for i := 0; i < totaldegoroutine; i++ {
//		go func() {
//			atomic.AddInt64(&contador, 1)
//			v := atomic.LoadInt64(&contador)
//			runtime.Gosched()
//			fmt.Println("Goroutines: ", v)
//			wg.Done()
//		}()
//	}

//	wg.Wait()
//	fmt.Println("Valor final: ", contador)
//}

///////////////////////////////////////////////////////
//EXERCICIO SEIS

package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("O seu sistema operacional é: ", runtime.GOOS)
	fmt.Println("A sua arquitetura é: ", runtime.GOARCH)

}
