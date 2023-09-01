/*package main

import (
	"testing"
)

// o nome da função de teste tem que começar nesse padrão Testnomepersonalizado
// esse argumento "(t *testing.T)" orbigatório também
// TestSoma, um teste para somar três números, onde todos são inteiros (int)
func TestSoma(t *testing.T) {
	teste := Soma(3, 2, 1)
	resultado := 6 //cria um teste a mão mesmo e vê se bate
	if teste != resultado {
		t.Error("Espected:", resultado, "Got:", teste) //se não bater exibe uma msg com essa cara
	}
}

// Esse falha
//func TestSoma2(t *testing.T) {
//	teste := Soma(3, 2, 1)
//	resultado := 7
//	if teste != resultado {
//		t.Error("Expected:", resultado, "Got:", teste)
//	}
//}

// TestMulti, um teste de multiplicação feito com dois números inteiros (int)
func TestMulti(t *testing.T) {
	teste := Multiplica(2, 2)
	resultado := 4
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

// TestPi, mais uma vez como o nome sugere é um teste para saber se aque é o valor de pi, mas dessa vez
// usando números com virgula (floatt64)
func TestPi(t *testing.T) {
	teste := Numpi(3.14)
	resultado := 3.14 //posso fazer assim também que da certo, sendo o mesmo valor ok
	if teste != resultado {
		t.Error("Expected:", resultado, "Got", teste)
	}
}

// TestNome, um teste para saber se o conteúdo da string está certo
func TestNome(t *testing.T) {
	teste := Nome("Pelo jeito pode string")
	//para dar certo com string o resultado precisa ser igual o teste ou vice e versa
	resultado := Nome("Pelo jeito pode string")
	if teste != resultado {
		t.Error("Expected:", resultado, "Got", teste)
	}
}

//O TESTE PODE SER FEITO DESSA FORMA TAMBÉM
//func TestNome(t *testing.T) {
//	teste := "Testando se assim dá tbm"
//	resultado := "Testando se assim dá tbm"
//	if teste != resultado {
//		t.Error("Expected", resultado, "Got:", teste)
//	}
//}
*/

////////////////////////////////////////////////////////////////////////
//TESTE EM TABELA

//Nesse tipo de teste fazemos muito mais rápido e sem precisar escrever tanto assim

/*package main

import "testing"

type test struct {
	data   []int //que serão os meus dados
	answer int   //que será os resultados
}

func TestSomaTabela(t *testing.T) {
	tests := []test{
		//a ordem dos valores digitados é: números a serem digitados e resultado certo
		test{data: []int{1, 2, 3}, answer: 6},
		//pode ser feito de maneira mais reduzida como o exemplo abaixo
		test{[]int{10, 11, 12}, 33},
		test{[]int{-5, 0, 5, 10}, 10},
	}
	for _, v := range tests {
		//nesse caso se eu coloco o soma..... na frente do if tenho que fazer o mesmo na frente do got
		//assim rodo ela só uma vez e não duas
		exemplo := Soma(v.data...)
		if exemplo != v.answer {
			t.Error("Expected:", v.answer, "Got", exemplo)
		}
	}
}

// PODE SER FEITO COM UM SLICE OF STRING E AI FICA DESSA FORMA:
type test2 struct {
	data   []string //tanto os dados quanto a resposta precisam ser um slice of string
	answer []string
}

func TestVariosTabela(t *testing.T) {
	testes := []test2{
		//nesse caso a string precisa ter a mesma quantidade da resposta/answer pra dar certo,
		//se tiver menor ou mais strings da erro
		test2{data: []string{"deu certo", "será que dá"}, answer: []string{"deu certo", "será que dá"}},
		test2{[]string{"repete isso", "repete de novo"}, []string{"repete isso", "repete de novo"}},
		test2{[]string{"mais uma vez", "teste"}, []string{"mais uma vez", "teste"}},
	}
	for _, v := range testes {
		exemplo := Nome(v.data...)
		//o v.answer precisa ser desse jeito pra dar certo com slice of string
		if exemplo != Nome(v.answer...) {
			t.Error("Expected:", v.answer, "Got", exemplo)
		}
	}
}

type test3 struct {
	data   []string
	//como esse teste é mais simples não existe necessidade de fazer um slice of string pro answer
	answer string
}

//ESSA É MAIS SIMPLES, SEM AQUELE MONTE DE STRING
func TestSimpTabela(t *testing.T) {
	teste := []test3{
		test3{data: []string{"Mais simples"}, answer: "Mais simples"},
		test3{[]string{"Só uma string"}, "Só uma string"},
		test3{[]string{"Sem segredo"}, "Sem segredo"},
	}
	for _, v := range teste {
		exemplo := Simplao(v.data...)
		//sem o monte de string posso usar apenas assim o v.answer
		if exemplo != v.answer {
			t.Error("Expected:", v.answer, "Got", exemplo)
		}
	}

}

func TestMulti(t *testing.T) {
	teste := Multiplica(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}*/

////////////////////////////////////////////////////////////////////////
//	TESTE SEM COISAS ESCRITAS, APENAS PARA VISUALIZAR

/*package main

import (
	"testing"
)

type test struct {
	data   []int
	answer int
}

func TestSomaTabela(t *testing.T) {
	tests := []test{
		test{data: []int{1, 2, 3}, answer: 6},
		test{[]int{10, 11, 12}, 33},
		test{[]int{-5, 0, 5, 10}, 10},
	}
	for _, v := range tests {
		exemplo := Soma(v.data...)
		if exemplo != v.answer {
			t.Error("Expected:", v.answer, "Got:", exemplo)
		}
	}
}

type test2 struct {
	data   []string
	answer []string
}

func TestVariosTabela(t *testing.T) {
	testes := []test2{
		test2{data: []string{"Será que da?", "Deu certo"}, answer: []string{"Será que da?", "Deu certo"}},
		test2{[]string{"Funciona bem", "assim também"}, []string{"Funciona bem", "assim também"}},
		test2{[]string{"Mais um teste", "cansei"}, []string{"Mais um teste", "cansei"}},
	}
	for _, v := range testes {
		exemplo := Nome(v.data...)
		if exemplo != Nome(v.answer...) {
			t.Error("Expected:", v.answer, "Got:", exemplo)
		}
	}
}

type test3 struct {
	data   []string
	answer string
}

func TestSimpTabela(t *testing.T) {
	teste := []test3{
		test3{data: []string{"Mais simples"}, answer: "Mais simples"},
		test3{[]string{"Só uma string"}, "Só uma string"},
		test3{[]string{"Sem segredo"}, "Sem segredo"},
	}
	for _, v := range teste {
		exemplo := Simplao(v.data...)
		if exemplo != v.answer {
			t.Error("Expected:", v.answer, "Got", exemplo)
		}
	}
}

func TestMulti(t *testing.T) {
	teste := Multiplica(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Exepecte:", resultado, "Got:", teste)
	}
}
*/

////////////////////////////////////////////////////////////////////////
//TESTES COMO EXEMPLOS

package main

import (
	"fmt"
	"testing"
)

// Sempre que for usar esse coloca "Example" antes
func ExampleSoma() {
	fmt.Println(Soma(3, 2, 1))
	//esse output vai funcionar como minha resposta/answer
	// Output: 6
}

func ExampleNome() {
	fmt.Println(Nome("Evangeline"))
	//Output: Evangeline
}

func ExampleMultiplica() {
	fmt.Println(Multiplica(5, 5))
	fmt.Println(Multiplica(10, 10))
	fmt.Println(Multiplica(100, 100))
	//Output: 25
	//100
	//10000
}

////////////////////////////////////////////////////////////////////////
//BENCHMARKS

//comando no terminal: go test -bench Nomedafunc

func BenchmarkMultiplica(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiplica(2, 2)
	}
}

func BenchmarkNome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Nome("Vendo se dá")
	}
}

type test struct {
	data   []int
	answer int
}

func BenchmarkSoma(b *testing.B) {
	tests := []test{
		test{data: []int{1, 2, 3}, answer: 6},
		test{[]int{10, 11, 12}, 33},
		test{[]int{-5, 0, 5, 10}, 10},
	}
	for i := 0; i < b.N; i++ {
		for _, v := range tests {
			Soma(v.data...)
		}
	}
}
