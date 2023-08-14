package main

import "fmt"

func main() {

	a := i()
	fmt.Println(a())
	fmt.Println(a())

	b := i()
	fmt.Println(b())
	fmt.Println(b())
}

func i() func() int {
	x := 0
	return func() int {  //pega essa funcionalidade e aplica do 0 no b
		x++                // por isso os números recomeçam na contagem e não seguem do a
		return x
	}
}

//tela
//1
//2
//1
//2
