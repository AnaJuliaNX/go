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
	fmt.Println("numero de goroutines 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("funcionando", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("erro checado 2:", ctx.Err())
	fmt.Println("numero de goroutines 2:", runtime.NumGoroutine())

	fmt.Println("prestes a cancelar o contexto")
	cancel()
	fmt.Println("contexto cancelado")

	time.Sleep(time.Second * 2)
	fmt.Println("erro checado 3:", ctx.Err())
	fmt.Println("numero de goroutines 3:", runtime.NumGoroutine())
}
