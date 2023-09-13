package usuario

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	//assim eu vinculo meu banco, lembrando que preciso no terminal dar um "go get"
	"github.com/AnaJuliaNX/cruid_basico/banco"
	"github.com/gorilla/mux"
)

type usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	//Aqui vou ler o corpo da requisição, ou seja, tudo que está no body
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuario para struct"))
		return
	}

	//Para fazer a conexão com o banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro de conexão com o banco de dados"))
		return
	}
	defer db.Close()

	/*
		Como inserir os dados no banco:
		-No GO usa o PREPARE STATEMENT que cria um comando para inserir e evitar o ataque SQL Injection
		O ponto de interrogação é usado pq de inicio não passei nenhum valor fixo e evita o ataque
		Nome e email ali são as colunas que eu criei no banco
	*/
	statement, erro := db.Prepare("Insert into usuarios(nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	//Aqui eu altero os pontos de interrogação para o nome e email digitados
	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	//Para devolver o Id do cara inserido no banco e para isso pego o Id primeiro.
	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o ID inserido"))
		return
	}

	//Status code, indica o que aconteceu com a requisição, se ela foi um sucesso ou falhou e etc
	w.WriteHeader(http.StatusCreated)
	//Aqui escrevo o Id inserido com um write
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso, ID: %d", idInserido)))
}

// BuscarUsuarios traz todos os usuarios salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	//Para abrir a conexão com o banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	//Para consultar as linhas da tabela com os usuários
	//uso um Query e dentro dele passo um "select * form usuarios"
	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários"))
		return
	}
	defer linhas.Close()

	//essa var é um slice of usuarios
	var usuarios []usuario
	for linhas.Next() {
		//essa é uma var simples, ele pegha cada usuário e vai enchendo o slice de usuários
		var usuario usuario
		//vou escanear a linha e jogar tudo para dentro dos dados do usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
		usuarios = append(usuarios, usuario)
	}
	//aqui eu transformo meu slice of usuarios em JSON
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuarios para JSON"))
		return
	}
}

// BuscarUmUsuario traz um usuário especifico no banco de dados
func BuscarUmUsuario(w http.ResponseWriter, r *http.Request) {

	//Leio o parametro que ta sendo passado
	parametros := mux.Vars(r)

	/*o parametro ID vai vir como string e converto ele para uint, uso o pacote strconv
		 Ele recebe três parametros:
	     1- o parametro["ID"],
		 2- a base do numero, nesse caso é 10 porque ta em decimal
		 3- o tamanho dos bits, pode ser um de 32, 64
	*/
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	//Faz a conexão com o banco de dados, abro depois de já ter o ID
	//Se der erro no ID não perco tempo abrindo o banco
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	//Vou ser mais especifica, pego os dados de apenas um usuário
	//nesse caso o que tenha o id igual ao que está vindo
	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuário"))
		return
	}

	var usuario usuario
	//uso if e não o for porque é um dado só que eu quero
	if linha.Next() {
		//Escaneio a linha e jogar tudo para dentro dos dados do usuario
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuario"))
			return
		}
	}
	//Transformo em JSON
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para JSON"))
		return
	}
}

// AtualizarUsuario altera os dados de um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	//Leio o meu parametro
	parametros := mux.Vars(r)

	/*Converto o meu parametro de string para int no mesmo esquema do de cima
	Tenho três parametros:
	1-O parametro de idetificação {"id"}
	2-A base do número, como ta em decimal é 10
	3-O tamanho dos bits, pode ser tanto 32 quanto 64 que não faz muita diferença
	*/
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao coverter o parametro para inteiro"))
		return
	}

	//Leio o corpo da requisição
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição "))
	}

	//Recebo ele como como JSON e transformo em struct
	var usuario usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	//Faço a conexão com o banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	/*Sempre usa o statement para qualquer operação que não seja de consulta,
	se for para inserir, deletar ou atualizar usa o statement. Passamos para ele
	os dados que vai precisar executar, no caso, nome, email e id
	*/
	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	//O dado ignorado exibe informações como o id e a quantidade de linhas afetadas e não quero nenhuma delas
	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário"))
		return
	}

	//No atualizar ele não vai exibir nada então só coloco isso aqui mesmo
	w.WriteHeader(http.StatusNoContent)
}

// DeletarUsuario remove um usuário do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	/*Converto o meu parametro de string para inteiro
	uso três parametros:
	1-Parametro de indetificação "parametros["id"]"
	2-A base do número, como está em decimal é 10
	3-O tamanho dos bits, pode ser tanto 32 quanto 64
	*/
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	//Faço a conexão com o banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	//Precisa colocar o "where id = ?" para não acabar deletando todos os usuários de uma vez
	//Passo para ele o dados que quero deletar
	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	//executo ele usando a função exec
	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
