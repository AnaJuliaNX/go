//TESTE DE UNIDADE

/*package enderecos

import (
	"testing"
)

//Teste unitário, não é tão viável nem confiável chances de dar erro é bem grande
func TestTipoEndereco(t *testing.T) {
	enderecoTeste := "Rua dos Sofredores"
	tipoEnderecoEsperado := "Avenida"
	tipoEnderecoRecebido := TipoEndereco(enderecoTeste)
	if tipoEnderecoRecebido != tipoEnderecoEsperado {
		t.Error("Expected:", tipoEnderecoEsperado, "Got:", tipoEnderecoRecebido)
	}
} */

package enderecos

import "testing"

type cenarioTeste struct {
	enderecoTeste string
	esperado      string
}

func TestTipoEnderco(t *testing.T) {

	cenariosteste := []cenarioTeste{
		{"Rua dos Sofredores", "Rua"},
		{"Avenida dos Livros", "Avenida"},
		{"Conjunto todos Juntos", "Conjunto"},
		{"Rodovia Esquina Azul", "Rodovia"},
		{"Praça das cores", "Endereço inválido"},
		{" ", "Endereço inválido"},
	}
	for _, cenario := range cenariosteste {
		recebido := TipoEndereco(cenario.enderecoTeste)
		if recebido != cenario.esperado {
			t.Error("Expected:", cenario.esperado, "Got:", recebido)
		}
	}
}
