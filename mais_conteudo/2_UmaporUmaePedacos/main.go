package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	//Pode ser exibido um por um
	fmt.Println("Dia:", now.Day())
	fmt.Println("Mês:", now.Month())
	fmt.Println("Ano:", now.Year())
	fmt.Println("Hora:", now.Hour())
	fmt.Println("Minuto:", now.Minute())
	fmt.Println("Segundo:", now.Second())

	//Pode ser exibido apenas os que eu quero
	fmt.Println("\nDia, mês e ano:", now.Day(), now.Month(), now.Year())
	fmt.Println("Hora, minutos, segundos:", now.Hour(), now.Minute(), now.Second())

}
