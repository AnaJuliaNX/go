package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //aparece como erro mas ta funcionando tudo ok
)

func main() {
	//urlConexão sereve para fazer uma conexão com o banco de dados
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	//retorna dois valores, uma conexão com o banco "db" e um erro
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	//db.Ping vai pingar a conexão, testar se ta conectado e tudo ok
	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta")

	//Para retornar dados para mim, no caso, as linhas da tabela ou um erro
	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	//fecho as linhas para não deixar esse cara em aberto
	defer linhas.Close()

	fmt.Println(linhas)
}
