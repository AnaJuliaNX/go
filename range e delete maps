package main

import "fmt"

func main() {
	listatele := map[string]int{
		"casa":     111,
		"pizzaria": 222,
		"hospital": 333,
	}
	
   //duas variaveis pq uma mostra a o escrito e outra o número

	for mostra, valor := range listatele {

		fmt.Println(mostra, valor)  
	}
}

//tela
//pizzaria 222
//hospital 333
//casa 111

//supondo que use só a variavel mostra:
//pizzaria 
//hospital 
//casa 

//supondo que use só a variavel valor:
// 222
// 333
// 111

  
DELETANDO UMA FATIA DO MAPS

package main

import "fmt"

func main() {
	listatele := map[string]int{
		"casa":     111,
		"pizzaria": 222,
		"hospital": 333,
	}

	delete(listatele, "casa")
	fmt.Println(listatele)
}

//tela
//map[hospital:333 pizzaria:222]
