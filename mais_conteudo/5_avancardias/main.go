package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	avancaDias := now.Add(15 * 24 * time.Hour)
	fmt.Println("O dia atual é:", now.Format("02/01/2006 03:04:05 -07"))
	fmt.Println("Daqui 15 dias será:", avancaDias.Format("02/01/2006 03:04:05 -07"))
}
