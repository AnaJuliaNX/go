package main

import (
	"fmt"
	"time"
)

func main() {

	determinada := time.Date(2005, 03, 03, 23, 50, 0, 0, time.Local)
	fmt.Println("Formatação dos USA:", determinada)
	fmt.Println("\nFormatação BR:", determinada.Format("02/01/2006 03:04:05"))
}
