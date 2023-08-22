package main

import "fmt"

func main() {

	umaslice := []int{1, 2, 3, 4}
	duasslice := []int{9, 10, 11, 12}

	fmt.Println(umaslice)
	fmt.Println(duasslice)

	umaslice = append(umaslice, 5, 6, 7, 8) //adiciona "pedaços" novos
	fmt.Println(umaslice)

	umaslice = append(umaslice, duasslice...) //adiciona uma slice nova
	fmt.Println(umaslice)

}


// tela
//[1 2 3 4]
//[9 10 11 12]
//[1 2 3 4 5 6 7 8]
//[1 2 3 4 5 6 7 8 9 10 11 12]
