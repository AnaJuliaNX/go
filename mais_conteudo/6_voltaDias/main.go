package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	voltaDias := now.Add(-15 * 24 * time.Hour)

	fmt.Println("O dia atual é:", now.Format("02/01/2006 03:04:05 -07"))
	fmt.Println("A quinze dias atrás foi:", voltaDias.Format("02/01/2006 03:04:05 -07"))

}
