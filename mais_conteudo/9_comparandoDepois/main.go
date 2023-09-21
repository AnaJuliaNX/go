package main

import (
	"fmt"
	"time"
)

func main() {

	primeiraData := time.Date(2024, 01, 01, 10, 10, 0, 0, time.Local)
	segundaData := time.Now()

	//Atenção: a data de fora é sempre comparada com a de dentro
	//É como se perguntassem: a data de fora vem depois da data de dentro?
	fmt.Println("A primiera data veio depois da segunda?", primeiraData.After(segundaData))

	//A resposta é sempre um valor booleano, ou seja true or false

}
