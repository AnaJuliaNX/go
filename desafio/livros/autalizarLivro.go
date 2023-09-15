package livros

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/AnaJuliaNX/desafio/banco"
	"github.com/AnaJuliaNX/desafio/modelo"
	"github.com/gorilla/mux"
)

func AtualizarUmLivro(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		ParaTratarErros(w, "Erro ao converter o parametro para int", 422)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		ParaTratarErros(w, "Erro ao ler o corpo", 422)
		return
	}

	var livro modelo.Livro
	erro = json.Unmarshal(corpoRequisicao, &livro)
	if erro != nil {
		ParaTratarErros(w, "Erro ao converter json para struct", 422)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		ParaTratarErros(w, "Erro ao se conectar com banco de dados", 422)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("Update livros set titulo = ?, autor = ? where id = ?")
	if erro != nil {
		ParaTratarErros(w, "Erro ao criar o statement", 422)
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(livro.Titulo, livro.Autor, ID); erro != nil {
		ParaTratarErros(w, "Erro ao atualizar o livro", 422)
		return
	}

	ParaTratarErros(w, "Livro atualizado com sucesso", 200)
	return
}
