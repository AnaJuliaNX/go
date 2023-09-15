package livros

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/AnaJuliaNX/desafio/banco"
	"github.com/AnaJuliaNX/desafio/modelo"
)

func CriarLivro(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		ParaTratarErros(w, "Erro ao ler o corpo", 422)
		return
	}

	var livro modelo.Livro
	if erro = json.Unmarshal(corpoRequisicao, &livro); erro != nil {
		fmt.Println(erro, 1)
		ParaTratarErros(w, "Erro ao converter json para struct", 422)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		ParaTratarErros(w, "Erro se conectar com o banco", 422)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into livros(titulo, autor, estoque) values(?, ?, ?)")
	if erro != nil {
		ParaTratarErros(w, "Erro ao criar o statement", 422)
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(livro.Titulo, livro.Autor, livro.Estoque)
	if erro != nil {
		ParaTratarErros(w, "Erro ao executar o statement", 422)
		return
	}

	_, erro = insercao.LastInsertId()
	if erro != nil {
		ParaTratarErros(w, "Erro ao obter o ID inserido", 422)
		return
	}

	ParaTratarErros(w, "Livro adicionado com sucesso", 200)
	return
}
