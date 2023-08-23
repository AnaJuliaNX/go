package main

import "fmt"

func main() {
	listatele := map[string]int{
		"Cam":  111,
		"Evie": 222,
		"Leo":  333, // a minha ta indo por ordem alfabetica, acho que vai assim
		"Mel":  444,
		"Nick": 555,
	}

	listatele["Raven"] = 777 //assim adiciono mais um na lista
	fmt.Println(listatele)
	fmt.Println(listatele["Evie"])

// comms ok
	if será, ok := listatele["Ben"]; !ok {
		fmt.Println("Não listado!")
		
              //esse comando exibe o Ben, se não estiver listado mostra o recado, se estiver mostra o número dele
	} else { 
		fmt.Println(será)
	}
}


//Tela

// map[Cam:111 Evie:222 Leo:333 Mel:444 Nick:555 Raven:777]
//222
//Não listado
