//COM PRINTLN = ele funciona da forma que conhecemos e o resultado sai na tela

//package main

//import (
//	"fmt"
//	"os"
//)

//func main() {
//	_, err := os.Open("sem-arquivo.txt")
//	if err != nil { //se o erro não estiver vazio eu lido com ele
//		fmt.Println("Aconteceu um erro", err)  //vai exibir isso na tela junto com nome do erro
//	}
//}

///////////////////////////////////////////////////
//EXEMPLO DOIS
//o package log implementa um pacote simples parta log (logar as coisas)
//escreve para standard error e printa na tela a data e hora de casa mensagem

//package main

//import (
//	"log"
//	"os"
//)

//func main() {
//	_, err := os.Open("no-file.txt")
//	if err != nil {
//		log.Println("Aconteceu um erro", err) //a diferença principal entre o log e o fmt é que
//ele vai dar a data e o horario de quando ocorreu o erro
//	}
//}

///////////////////////////////////////////////////
//EXEMPLO TRÊS
//para conferir se o erro ta dentro do "log.txt" mesmo eu uso o comando no terminal "cat log.txt"
//vai exibir dentro do arquivo "log.txt" a data e hora do erro

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt") //cria um arquivo log
	if err != nil {                //verifica se tem erro e se tiver exibe o erro
		fmt.Println(err)
	}
	defer f.Close()  //depois que eu criei o arquivo eu fecho ele
	log.SetOutput(f) //isso quer dizer que se acontecer alguma coisa com o "log" vai salvar tudo aqui
	//e vai escrever no arquivo já que eu tenho um ponteiro para um arquivo

	f2, err := os.Open("no-file.txt") //aqui to tentando abrir um arquivo que não existe
	if err != nil {                   //obviamente vai dar erro
		//o erro vai ser salvo dentro do aquivo "log.txt"
		log.Println("Aconteceu um erro", err)
	}
	defer f2.Close()

	fmt.Println("Checando o aquivo log.txt no diretório")
}
