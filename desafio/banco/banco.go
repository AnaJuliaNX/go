package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	stringconexao := "dono:livraria@/livraria?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringconexao)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
