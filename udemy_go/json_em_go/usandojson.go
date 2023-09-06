package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type livros struct {
	Autor  string `json:"autor"`
	Titulo string `json:"titulo"`
	Nota   int    `json:"nota"`
}

func main() {

	livro := livros{"Rick Riordan", "PJO", 10}

	//primeiro ele vai converter nossos dados digitados em um slyce of bytes, ou seja, em  vários números
	livroEmJson, erro := json.Marshal(livro)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(livroEmJson)
	//bytes.NewBuffer, com isso vou receber o meu slyce of bytes e converter par JSON
	fmt.Println(bytes.NewBuffer(livroEmJson))

	//primeiro ele vai converter nossos dados digitados em um slyce of bytes, ou seja, em vários números
	livro2 := map[string]string{
		"autor": "Raphael Montes",
		"nome":  "Jantar secreto",
		"nota":  "9",
	}
	fmt.Println("Mensagem 2")

	livro2EmJson, erro := json.Marshal(livro2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(livro2EmJson)
	//bytes.NewBuffer, com isso vou receber o meu slyce of bytes e converter para JSON
	fmt.Println(bytes.NewBuffer(livro2EmJson))
}
