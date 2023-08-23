package main

import "fmt"

func main() {
	x := retorna()
	y := x(3)
	fmt.Println(y)    //nesse caso eu tenho uma função dentro de outra função
}
func retorna() func(int) int {
	return func(i int) int {     //nesse caso eu tenho uma função dentro de outra função
		return i * 10
	}
}
//tela
//30

///////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {

	t := numeroimpar(soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	fmt.Println(t)
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func numeroimpar(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {     //passa uma função como argumento de função
		if v%2 != 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}
