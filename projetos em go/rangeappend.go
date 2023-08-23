package main

import "fmt"

func main() {

	slice := []string{"arroz", "feijão", "farofa"}

	//slice[4] = "melancia"  Assim não funcina

	slice = append(slice, "e sobremesa") //assim pode

	for indice, valor := range slice {

		fmt.Println("No indice", indice, "tem:", valor)
	}
}


//também pode ser feito assim, mostrando só o que tem escrito no slice e não seu valor
// mas vale ressaltaqr que precisa ter uma variavel (???) de indice sem uso "_" para não aparecer os valores

package main

import "fmt"

func main() {

	slice := []string{"arroz", "feijão", "farofa"}

	//slice[4] = "melancia"  Assim não funcina

	slice = append(slice, "e sobremesa") //assim pode

	for _, valor := range slice {

		fmt.Println("Esse slice tem:", valor)
	}
}
