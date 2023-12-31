package main

import "fmt"

func main() { 

	var a = 10
	var ponteiro = &a

	fmt.Println("endereço do a:", &a) //é o endereço do valor armazenado em a
	fmt.Println("valor do a:", a)
	
	fmt.Println("endereço do ponteiro:", ponteiro) //é o endereço do valor do a
	fmt.Println("valor do ponteiro:", *ponteiro)  
//ao usar * pega o valor do ponteiro o valor do ponteiro é o valor armazenado no endereço do a
}

//tela
//endereço do a: 0xc000012028     //cada hora que atualiza é um endereço diferente
//valor do a: 10
//endereço do ponteiro: 0xc000012028  //cada hora que atualiza é um endereço diferente
//valor do ponteiro: 10


////////////////////////////////////////////////////////////////////

			EXEMPLO MAIS LIMPO
package main

import "fmt"

func main() {

	var valor = 10
	var ponteiro = &valor  //endereço do valor da váriavel

                                   //exibe:
	fmt.Println(valor)    //valor da variavel
	fmt.Println(ponteiro)  //endereço da variavel
	fmt.Println(*ponteiro)  //valor da variavel armazenada nesse endereço
}
//tela
//10
//0xc000012028
//10


//////////////////////////////////////////////////////////////////

	EXEMPLO QUE MUDA NA MEMÓRIA E NÃO UMA CÓPIA

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {

	aelin := pessoa{"Celaena", "Sardothien", 21}
	fmt.Println(aelin)

	mudeMe(&aelin)  //essa função alterou po original e não uma cópia
	fmt.Println(aelin)
}

func mudeMe(p *pessoa) {
	(*p).nome = "Aelin"
	(*p).sobrenome = "Muito dificl de digitar"
      // p.idade = 22 (pode ser feito dessa forma também)
}
//tela
//{Celaena Sardothien 21}
//{Aelin Muito dificl de digitar 21}
