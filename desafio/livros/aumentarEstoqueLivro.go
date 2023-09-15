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

func AumentarEstoque(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		ParaTratarErros(w, "Erro ao se conectar com o banco", 422)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		ParaTratarErros(w, "Erro ao se conectar com o banco", 422)
		return
	}

	linhas, erro := db.Query("select * from livros where id = ?", ID)
	if erro != nil {
		ParaTratarErros(w, "Erro ao buscar livros", 422)
		return
	}

	var livro modelo.Livro
	if linhas.Next() {
		if erro := linhas.Scan(&livro.ID, &livro.Titulo, &livro.Autor, &livro.Estoque); erro != nil {
			ParaTratarErros(w, "Erro ao escanear livros", 422)
			return
		}
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		ParaTratarErros(w, "Erro ao ler ao ler os dados da requisição", 422)
		return
	}

	var dadosRequest modelo.DadosDoRequestAdicionarEstoque
	erro = json.Unmarshal(corpoRequisicao, &dadosRequest)
	if erro != nil {
		ParaTratarErros(w, "Erro ao pegar dados inseridos", 422)
		return
	}

	if dadosRequest.Quantidade <= 0 {
		ParaTratarErros(w, "Quantidade inserida não pode ser menor do que zero", 422)
		return
	}

	livro.Estoque = livro.Estoque + dadosRequest.Quantidade

	erro = AlterarOEstoque(int(ID), livro.Estoque)
	if erro != nil {
		ParaTratarErros(w, erro.Error(), 422)
		return
	}
}
