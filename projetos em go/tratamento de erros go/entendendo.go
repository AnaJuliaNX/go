package main

import (
	"fmt"
)

func main() {
	var quest1, quest2, quest3 string

	fmt.Println("Nome: ")
	_, err := fmt.Scan(&quest1) //espera eu digitar um valor, no caso seria o usuário digitando
//depois ela retorna o numero de bytes e um erro, o bytes fica "_" porque não me interessa
	if err != nil {
		panic(err)
	} // se a minha variavel erro foi diferente de nada eu preciso lidar ccom o erro e pra isso uso o "panic"

	fmt.Println("Comida favorita: ")
	_, err = fmt.Scan(&quest2)
	if err != nil { //faz o mesmo que no de cima
		panic(err)
	}

	fmt.Println("Esporte favorito: ")
	_, err = fmt.Scan(&quest3) //faz o mesmo que no primeiro
	if err != nil {
		panic(err)
	}

	fmt.Println(quest1, "gosta de", quest2, "e também de", quest3)
//só exibe as minhas respostas salvas no scan
}

///////////////////////////////////////////////////
//EXEMPLO DOIS

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("nome.txt") //isso é onde posso escrever
	if err != nil {
		fmt.Println(err) //esse erro no caso vai exbir na tela o nome do erro
		return           //depois de mostrar o nome do erro ele retorna pro começo da função e ela vai encerrar
	}
	defer f.Close()
	defer fmt.Println("Rodou tudo certo aqui")

	r := strings.NewReader("Gato de Botas") //isso é onde posso ler

	io.Copy(f, r) //essa pega uma coisa e coloca em outra, no caso, pega o negocio que posso ler e ecoloca
//onde posso escrever
}

///////////////////////////////////////////////////
//EXEMPLO TRÊS

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("nomes.txt") //vou abrir esse arquivo
	if err != nil {
		fmt.Println(err) //se tiver algum nerro ele vai me reternornar o erro e encerro
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f) //assim que eu abri o aerquivo eu vou ler ele com essa função
	if err != nil {
		fmt.Println(err) //me retorna um erro caso exista algum erro
		return
	}
	fmt.Println(string(bs)) //se der tudo certo ele vai executar o arquivo.
	//no caso, se ele abriiu, leu e não deu nada errado então ele exibe na tela
}
