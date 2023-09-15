package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AnaJuliaNX/desafio/livros"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/livros", livros.CriarLivro).Methods(http.MethodPost)
	router.HandleFunc("/livros", livros.BuscarLivros).Methods(http.MethodGet)
	router.HandleFunc("/livros/{id}", livros.BuscarUmLivro).Methods(http.MethodGet)
	router.HandleFunc("/livros/{id}", livros.AtualizarUmLivro).Methods(http.MethodPut)
	router.HandleFunc("/livros/{id}/diminuir", livros.DiminuirEstoqueLivro).Methods(http.MethodPut)
	router.HandleFunc("/livros/{id}/aumentar", livros.AumentarEstoque).Methods(http.MethodPut)
	router.HandleFunc("/livros/{id}", livros.DeletarUmLivro).Methods(http.MethodDelete)

	fmt.Println("Executando na porta 4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}
