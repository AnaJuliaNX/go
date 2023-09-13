package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AnaJuliaNX/cruid_basico/usuario"

	"github.com/gorilla/mux"
)

/*
CRUID: Created, Read, Update, Delete (isso que a sigla significa)

Created = Post
Read = Get
Update = Put
Delete = Delete
*/
func main() {
	router := mux.NewRouter()
	//isso quer dizer que quando a minha rota de usuários receber um POST, ela vai executar a func
	//MethodPost cria um usuário novo
	router.HandleFunc("/usuarios", usuario.CriarUsuario).Methods(http.MethodPost)
	//MethodGet busca os usuários
	router.HandleFunc("/usuarios", usuario.BuscarUsuarios).Methods(http.MethodGet)
	//MethodGet com um /{id} busca um usuário especifico, usa as chaves "{}" para que
	//o go entenda que dentro é um valor variádico e não fixo
	router.HandleFunc("/usuarios/{id}", usuario.BuscarUmUsuario).Methods(http.MethodGet)
	//MethodPut com um /{id} atualiza os dados de um usuário especifico
	router.HandleFunc("/usuarios/{id}", usuario.AtualizarUsuario).Methods(http.MethodPut)
	//MethodDelete com um /{id} deleta um usuário especifico
	router.HandleFunc("/usuarios/{id}", usuario.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
