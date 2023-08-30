//EXERCICIO UM

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Testando se funciona", "já que antes deu erro", "aaaaaa"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(bs))
}

////////////////////////////////////////////////////////////////
//EXERCICIO DOIS

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type estoque struct {
	Produto    string
	Quantidade int
	Descrição  []string
}

func main() {
	p1 := estoque{
		Produto:    "A revolução dos bichos",
		Quantidade: 14,
		Descrição:  []string{"Livro bruchura", "300 páginas", "Distopia"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln("Produto não encontrado")
	}

	fmt.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) { //nesse caso estamos retornando um sliceofbytes e um erro
	bs, err := json.Marshal(a)
	if err != nil {   //aqui eu defino uma função
		return []byte{}, fmt.Errorf("Produto não encontrado")
	}
	return bs, nil
}

////////////////////////////////////////////////////////////////
//EXERCICIO TRÊS

package main

import (
	"fmt"
)

type erroEspecial struct {
	teste string
}

func (e erroEspecial) Error() string { //preciso implementar uma interface
	return fmt.Sprintf("Deu errado, olha só: %v", e.teste)
}

func erroComParametro(e error) {
	fmt.Println("Passado como argumento:", e.(erroEspecial).teste, "\nNo método Error eu tenho:", e)
}
func main() {
	arg := erroEspecial{"TESTANDO ISSO"}
	erroComParametro(arg)
}

////////////////////////////////////////////////////////////////
//EXERCICIO QUATRO

package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	num1 string
	num2 string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("Erro de calculo: %v %v %v", se.num1, se.num2, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//meu jeito simplão
		//return 0, fmt.Errorf("Calculo inválido: o número %v não pode ser negativo", f)

		//jeito da moça
		erroNovo := fmt.Errorf("Deu erro no valor: %v", f)
		return 0, sqrtError{"3.14 e", "7", erroNovo}
	}
	return 42, nil
}

