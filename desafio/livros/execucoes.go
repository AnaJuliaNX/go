package livros

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/AnaJuliaNX/desafio/modelo"

	"github.com/AnaJuliaNX/desafio/banco"
)

var livro modelo.Livro
var dadosDoRequestDiminuirEstoque modelo.DadosDoRequestDiminuirEstoque
var dadosDoRequestAdicionarEstoque modelo.DadosDoRequestAdicionarEstoque

func ParaTratarErros(w http.ResponseWriter, message string, statusCode int) {

	w.WriteHeader(statusCode)
	w.Write([]byte(message))

}
func ConectarOBanaco() (*sql.DB, error) {
	db, erro := banco.Conectar()
	if erro != nil {
		return nil, errors.New("erro ao se conectar com banco")
	}
	return db, nil
}
func ParaBuscarOsLivros() ([]modelo.Livro, error) {

	db, erro := ConectarOBanaco()
	if erro != nil {
		return nil, erro
	}
	defer db.Close()

	linhas, erro := db.Query("select * from livros")
	if erro != nil {
		return nil, errors.New("erro ao buscar livros")
	}
	defer linhas.Close()

	var livros []modelo.Livro
	for linhas.Next() {
		var livro modelo.Livro

		if erro := linhas.Scan(&livro.ID, &livro.Titulo, &livro.Autor, &livro.Estoque); erro != nil {
			return nil, errors.New("erro ao escanear livros")
		}

		livros = append(livros, livro)
	}
	return livros, nil
}

func ParaBuscarUMlivro(ID int) (modelo.Livro, error) {

	db, erro := ConectarOBanaco()
	if erro != nil {
		return modelo.Livro{}, erro
	}
	defer db.Close()

	linhas, erro := db.Query("select * from livros where id = ?", ID)
	if erro != nil {
		return modelo.Livro{}, errors.New("erro ao buscar livros")
	}
	defer linhas.Close()

	var livroencontrado modelo.Livro
	if linhas.Next() {
		erro := linhas.Scan(&livroencontrado.ID, &livroencontrado.Titulo, &livroencontrado.Autor, &livroencontrado.Estoque)
		if erro != nil {
			return modelo.Livro{}, errors.New("erro ao buscar livros")
		}
	}
	return livroencontrado, nil
}

func AlterarOEstoque(ID int, estoque int) error {

	db, erro := ConectarOBanaco()
	if erro != nil {
		return erro
	}
	defer db.Close()

	statement, erro := db.Prepare("Update livros set estoque = ? where id = ?")
	if erro != nil {
		return errors.New("erro ao buscar livros")
	}
	defer statement.Close()

	if _, erro := statement.Exec(estoque, ID); erro != nil {
		return errors.New("erro ao buscar livros")
	}
	return nil
}
