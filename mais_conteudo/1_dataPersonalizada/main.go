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

	//ATENÇÃO: A data de fora é sempre comparada com a de dentro
	//Comparando se uma data veio antes da outra:
	fmt.Println("Verificando se veio antes:", menosDias.Before(maisDias))
	//Comparando se uma data é depois da outra
	fmt.Println("Verificando se veio depois:", menosDias.After(maisDias))

}
