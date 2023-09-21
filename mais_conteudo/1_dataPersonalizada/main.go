package main

import (
	"fmt"
	"time"
)

func main() {

	// Horário local
	now := time.Now()
	fmt.Println("Formato padrão do Go:", now)
	//Formatando como usamos aqui no Brasil
	fmt.Println("\nFormato personalizado:", now.Format("02/01/2006, 03:04:05"))
}
