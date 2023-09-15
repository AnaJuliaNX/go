package livros

import (
	"encoding/json"
	"net/http"
)

func BuscarLivros(w http.ResponseWriter, r *http.Request) {

	livros, erro := ParaBuscarOsLivros()
	if erro != nil {
		ParaTratarErros(w, erro.Error(), 422)
		return
	}

	erro = json.NewEncoder(w).Encode(livros)
	if erro != nil {
		ParaTratarErros(w, "Erro ao converter para json", 422)
		return
	}

	ParaTratarErros(w, "Livros buscados com sucesso", 200)
	return
}
