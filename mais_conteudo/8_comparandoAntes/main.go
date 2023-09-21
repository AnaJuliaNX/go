package main

import (
	"fmt"
	"time"
)

func main() {

	primeiraData := time.Now()
	segundaData := time.Date(2005, 03, 03, 23, 50, 0, 0, time.Local)

	//Atenção: a data de fora é sempre comparada com a de dentro
	//É como se perguntassem: a data de fora vem antes da data de dentro?
	fmt.Println("A primeira data veio antes da segunda?", primeiraData.Before(segundaData))

	//O resultado da comparação é um valor booleano, ou seja, true or false
}
