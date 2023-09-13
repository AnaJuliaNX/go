package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver de coenxão com o Mysql
)

// Esse arquivo vai servir para mim fazer a conexão com o banco de dados
// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {

	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		//nesse caso se der um erro não vou ter como retornar nada, então retorno um valor vazio e um erro
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil //nesse caso vou retornar o db e nada, já que não tem erro
}

//OBS: em go posso ter qyantos retornos achar necessário que não da BO, mas o recomendável é três
