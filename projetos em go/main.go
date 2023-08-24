package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Esse é o teste de compilação cruzada, foi compilado em um linux e agora está em um sistema: ",
		runtime.GOARCH, runtime.GOOS)

}

//Isso é util quando quero compular algo, por exemplo, de um Linux para um windows ou mac
//GOARCH = se refere a aqrquitetura do sistma
//GOOS = se refere ao sistema operacional
//exmplo: GOARCH: Linux GOOS: arm64

//NO TERMINAL:
//GOOS=windows GOARCH=amd64 go build main.go  (para windows)
//GOOS=darwin GOARCH=amd64 go build main.go  (para mac)
//esse main.go no caso é o meu arquivo onde tá o código que vai me permitir fazer isso, o nome muda de acordo
//com o seu arquivo
//Caso queira conferir é só dar um "ls" no terminal e ver se tem um arquivo com o mesmo nome do seu mas
//o final sendo .exe para windows e nada para mac
