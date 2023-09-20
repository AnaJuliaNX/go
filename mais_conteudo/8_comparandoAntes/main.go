package main

import (
	"fmt"
	"time"
)

func main() {

	primeiraData := time.Now()
	segundaData := time.Date(2005, 03, 03, 23, 50, 0, 0, time.Local)

	fmt.Println("A primeira data veio antes da segunda?", primeiraData.Before(segundaData))
}
