package main

import "fmt"

func main() {

	slice := []int{1, 2, 6, 2, 3, 5, 7, 7, 4}
	var slicefila []int

	//fmt.Println(slice[:2], slice[3:4], slice[6:])

	for i := 0; i < len(slice); i++ {
		if i == 0 || slice[i] >= slice[i-1] {
			if i > 0 {
				fmt.Println(slice[i], slice[i-1])
			}
			slicefila = append(slicefila, slice[i])
		}
	}
	fmt.Println(slicefila)
}

//if i <= 2 {
//fmt.Println(slice[i])
//} else if i >= 6 {
//fmt.Println(slice[i])
//}
