package livros

import (
	"net/http"
	"strconv"

	"github.com/AnaJuliaNX/desafio/banco"
	"github.com/gorilla/mux"
)

func DeletarUmLivro(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		ParaTratarErros(w, "Erro ao converter o parametro para int", 422)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		ParaTratarErros(w, "Erro ao se conectar com o banco", 422)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from livros where id = ?")
	if erro != nil {
		ParaTratarErros(w, "Erro ao criar o statement", 422)
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		ParaTratarErros(w, "Erro ao deletar o livro", 422)
		return
	}

	ParaTratarErros(w, "Livro deletado com sucesso", 200)
	return
}
