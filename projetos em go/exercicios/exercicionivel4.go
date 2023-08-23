package main

import "fmt"

func main() {
	//             0  1  2  3  4  5  6  7  8  9  <-estão assim
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice[:3])               //primeiro ao terceiro item
	fmt.Println(slice[4:])               //quinto ao ultimo slice
	fmt.Println(slice[1:7])              //segundo ao sétimo slice
	fmt.Println(slice[2:9])              //terceito ao penultimo slice
	fmt.Println(slice[8 : len(slice)-1]) //apenas o penultimo slice

}

//tela 
//[1 2 3]
//[5 6 7 8 9 10]
//[2 3 4 5 6 7]
//[3 4 5 6 7 8 9]
//[9]


////////////////////////////////////////////


package main

import "fmt"

func main() {
	mapa := map[string][]string{
		"verstappen_max": {"ser chato", "ganhar"},
		"senna_ayrton":   {"brasileiro", "ser o melhor"},
	}

	mapa["leclerc_charles"] = []string{"correr", "bater"}

	delete(mapa, "verstappen_max")

	for k, v := range mapa {
		fmt.Println(k)

		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}
}

//tela
//senna_ayrton
//	 0  -  correr
//	 1  -  ganhar
//leclerc_charles
//	 0  -  correr
//	 1  -  bater	
//OBS: Max não aparece porque deletei ele
