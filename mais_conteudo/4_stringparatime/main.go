package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	datasalva := "03/03/2005 23:50:10 -03"
	parsed, erro := time.Parse("02/01/2006 15:04:05 -07", datasalva)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println("Data selecionada:", parsed.Format("02/01/2006 15:04:05 -07"))
}
