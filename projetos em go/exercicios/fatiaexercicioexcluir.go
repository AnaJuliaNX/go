package main

import "fmt"

func main() {

	sabores := []string{"frango,", "bacon,", "calabresa,", "queijo"}

	fatia := sabores[1:4]

	fmt.Println(fatia)
}


// tela =

//  [bacon, calabresa, queijo]



package main

import "fmt"

func main() {

	sabores := []string{"frango,", "bacon,", "calabresa,", "queijo"}

	fatia := sabores[2:len(sabores)]

	//len tem a função de exibir todo o resto da string

	fmt.Println(fatia)

	//pode ser feito dessa forma também para exibir tudo

	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}
}


// tela =

// frango,
// bacon, 
// casalabresa,
// queijo



//como apagar uma fatia do slice
package main

import "fmt"

func main() {
                                             //vou apagar essa
	sabores := []string{"frango,", "bacon,", "abacaxi", "calabresa,", "queijo"}

	sabores = append(sabores[:2], sabores[3:]...)
	fmt.Println(sabores)
}

// tela
//[frango, bacon, calabresa, queijo]
