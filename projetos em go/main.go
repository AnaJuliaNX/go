package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Esse é o teste de compilação cruzada, foi compilado em um linux e agora está em um sistema: ",
		runtime.GOARCH, runtime.GOOS)

}
