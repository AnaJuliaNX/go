package main

import (
	"fmt"
	"time"
)

func main() {

	primeiraData := time.Now()

	segundaData := time.Date(2023, 10, 23, 10, 50, 0, 0, time.Local)

	fmt.Println("AS datas são compativeis?", primeiraData.Equal(segundaData))

	//O resultado da comparação é um valor booleano, ou seja, apenas true ou false
}
