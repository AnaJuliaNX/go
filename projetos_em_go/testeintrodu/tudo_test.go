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

package main

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

func TestSoma(t *testing.T) {
	teste := Soma(2, 3, 5)
	resultado := 10
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

func TestMulti(t *testing.T) {
	teste := Multiplica(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}
