package enderecos

import "strings"

// Exercicio de teste, para estudar sobre testes
func TipoEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Rodovia", "Conjunto"}

	//essa ação vai fazer trasformar a frase em um slice of string e pegar apenas o primeiro para mim
	//ficaria assim: "Rua dos Livros", que vira uma slice: ["Rua", "dos", "Livros"] e ai ele me retorna
	// o de posição 0, que nesse caso é a "Rua"
	// esse " " na frente do endereco é o indicador onde começará e onde terminará a slice, poderia ser
	//definido por até por uma letra, mas nesse caso é por um espaço em branco
	primeiraPalavraEndereco := strings.Split(endereco, " ")[0]

	enderecoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoValido = true
		}
	}

	if enderecoValido {
		return primeiraPalavraEndereco
	}
	return "Tipo inválido ou escrito com letra minuscula"
}
