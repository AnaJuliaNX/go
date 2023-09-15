package livros

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/AnaJuliaNX/desafio/modelo"
	"github.com/gorilla/mux"
)

func DiminuirEstoqueLivro(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		ParaTratarErros(w, "Erro ao converter o parametro para inteiro", 422)
		return
	}

	livrobuscado, erro := ParaBuscarUMlivro(int(ID))
	if erro != nil {
		ParaTratarErros(w, "Erro ao converter o parametro para int", 422)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		ParaTratarErros(w, "Erro ao ler dados do corpo", 422)
		return
	}

	var dadosRequest modelo.DadosDoRequestDiminuirEstoque
	erro = json.Unmarshal(corpoRequisicao, &dadosRequest)
	if erro != nil {
		ParaTratarErros(w, "Erro ao converter de json para struct", 422)
		return
	}

	if livrobuscado.Estoque > dadosRequest.Quantidade {
		livrobuscado.Estoque = livrobuscado.Estoque - dadosRequest.Quantidade

	} else {
		fmt.Println("Estoque insuficiente")
	}

	erro = AlterarOEstoque(int(ID), livrobuscado.Estoque)
	if erro != nil {
		ParaTratarErros(w, erro.Error(), 422)
		return
	}
}
