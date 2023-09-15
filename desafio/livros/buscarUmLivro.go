package livros

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func BuscarUmLivro(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		ParaTratarErros(w, "Erro ao converter o parametro para int", 422)
		return
	}

	livrobuscado, erro := ParaBuscarUMlivro(int(ID))
	if erro != nil {
		ParaTratarErros(w, "Erro ao converter o parametro para int", 422)
		return
	}

	erro = json.NewEncoder(w).Encode(livrobuscado)
	if erro != nil {
		ParaTratarErros(w, "Erro ao converter para json", 422)
		return
	}

	ParaTratarErros(w, "Livro buscado com sucesso", 200)
	return
}
