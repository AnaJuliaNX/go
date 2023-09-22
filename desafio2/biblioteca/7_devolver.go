package biblioteca

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/AnaJuliaNX/desafio2/dados"
	"github.com/gorilla/mux"
)

func Devolvendo(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	//Busco o usuário pelo ID
	usuario_id, erro := strconv.ParseUint(parametro["usuario_id"], 10, 32)
	if erro != nil {
		fmt.Println(erro, 1)
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
		return
	}

	usuariobuscado, erro := buscandoUMUsuario(int(usuario_id))
	if erro != nil {
		TratandoErros(w, "Erro ao buscar o usuário", 422)
		return
	}

	//Busco o livro pelo ID
	livro_id, erro := strconv.ParseUint(parametro["livro_id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
		return
	}

	livrobuscado, erro := BuscandoUMLivro(int(livro_id))
	if erro != nil {
		TratandoErros(w, "Erro ao buscar o livro", 422)
		return
	}

	db, erro := ConectandoNoBanco()
	if erro != nil {
		TratandoErros(w, "Erro ao se conectar com o banco de dados", 422)
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select nome_usuario, titulo_livro, data_emprestimo, data_devolucao, taxa_emprestimo from emprestimo_devolucao where nome_usuario = ? and titulo_livro = ?", usuariobuscado.Nome, livrobuscado.Titulo)
	if erro != nil {
		TratandoErros(w, "Erro ao buscar dados dos emprestimos", 422)
		return
	}
	defer linhas.Close()

	var emprestado dados.EmprestimoDevolucao
	if linhas.Next() {
		err := linhas.Scan(
			&emprestado.Nome_Usuario,
			&emprestado.Titulo_livro,
			&emprestado.Data_Emprestimo,
			&emprestado.Data_Devolucao,
			&emprestado.Taxa_Emprestimo,
		)

		if err != nil {
			TratandoErros(w, "Erro ao escanear os dados dos empréstimos", 422)
			return
		}
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		TratandoErros(w, "Erro ao ler os dados do corpo", 422)
		return
	}

	var devolver dados.EmprestimoDevolucao
	erro = json.Unmarshal(corpoRequisicao, &devolver)
	if erro != nil {
		TratandoErros(w, "Erro ao converter json para struct", 422)
		return
	}

	//Confiro se o usuário da devolução é o mesmo que fez emprestimo
	if usuariobuscado.Nome != emprestado.Nome_Usuario {
		fmt.Println(erro, 1)
		fmt.Println("As credenciais do usuário são incompativeis")
	}

	//Confiro se o livro a ser devolvido é o mesmo que foi emprestado
	if livrobuscado.Titulo != emprestado.Titulo_livro {
		fmt.Println(erro, 1)
		fmt.Println("Titulo do livro diferente do emprestado")
	}

	//devolucaoHoje := time.Date(2023, 10, 21, 20, 10, 0, 0, time.Local)

	//Data da devolução em tempo real
	devolucaoHoje := time.Now()

	//Formato para que consiga fazer a comparação das datas
	hojeFormatada := devolucaoHoje.Unix()
	devolucaoFormatada := emprestado.Data_Devolucao.Unix()

	if hojeFormatada < devolucaoFormatada {
		fmt.Println("Usuário:", emprestado.Nome_Usuario)
		fmt.Println("Livro:", emprestado.Titulo_livro)
		fmt.Println("Devolução feita antes do prazo de 15 dias")

		livrobuscado.Estoque = emprestado.Quantidade + livrobuscado.Estoque

	} else {
		fmt.Println("Usuário:", emprestado.Nome_Usuario)
		fmt.Println("Livro:", emprestado.Titulo_livro)
		fmt.Println("Devolução feita após o prazo de quinze dias")
		fmt.Println("Multa de: R$", emprestado.Taxa_Emprestimo)

		livrobuscado.Estoque = emprestado.Quantidade + livrobuscado.Estoque
	}

}
