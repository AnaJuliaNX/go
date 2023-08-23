                           ORDENANDO TUDO
                            MEU EXEMPLO
package main

import (
	"fmt"
	"sort"
)

type livros struct {
	Nome       string
	Personagem string
	Idade      int
	Nomeperso  []string  //SLICE OF STRING PARA PODE COLOCAR MAIS DE UM ELEMENTO
}

type porIdade []livros

func (p porIdade) Len() int           { return len(p) }  //RETORNA OS ELEMENTOS
func (p porIdade) Less(i, j int) bool { return p[i].Idade < p[j].Idade } 
//RETORNA SE UM É MAIS NOVO(MENOR) QUE O OUTRO E VICE-VERSA
func (p porIdade) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
//TROCA A ORDEM DOS ELEMENTOS PARA QUE FIQUEM DE ACORDO COM O LESS

type porLivro []livros

func (p porLivro) Len() int           { return len(p) } //RETORNA OS ELEMENTOS
func (p porLivro) Less(i, j int) bool { return p[i].Nome < p[j].Nome } 
//RETORNA SE UM É MAIS NOVO(MENOR) QUE O OUTRO E VICE-VERSA
func (p porLivro) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
//TROCA A ORDEM DOS ELEMENTOS PARA QUE FIQUEM DE ACORDO COM O LESS

func main() {

	p1 := livros{
		Nome:       "Trono de vidro",
		Personagem: "Aelin,",
		Idade:      21,
		Nomeperso: []string{  //SLICE OF STRING
			"Celaena",
			"Matadora de bruxas",
			"Rainha do fogo",
		},
	}
	p2 := livros{
		Nome:       "Percy Jackason",
		Personagem: "Percy,",
		Idade:      18,
		Nomeperso: []string{  //SLICE OF STRING
			"Perseus Jackson",
			"Peter Jhonson ",
			"cabeça de alga",
		},
	}
	p3 := livros{
		Nome:       "Devils Night",
		Personagem: "Winter,",
		Idade:      25,
		Nomeperso: []string{  //SLICE OF STRING
			"Cegueta",
			"Litlle Monster",
			"Inverno",
		},
	}

	personagens := []livros{p1, p2, p3}
	fmt.Println(personagens)  //EXIBE TODOS NA ORDEM QUE FORAM DIGITADOS

	for _, v := range personagens {
		sort.Strings(v.Nomeperso)  //COLOCA OS "APELIDOS" EM ORDEM ALFABÉTICA
  //RESSALTANDO QUE ACENTOS E LETRAS MINUSCULAS, INDEPENDENTE DA LETRA, VEM DEPOIS DAS MAISCULAS
	}

	sort.Sort(porIdade(personagens))
	fmt.Println("Ordena por idade:\n")  //EXIBE EM ORDEM ALFABÉTICA CRESCENTE OU DECRESCENTE, DEPENDE DO QUE EU QUISER
	harmonia(personagens)

	sort.Sort(porLivro(personagens))
	fmt.Println("Ordena por nome do livro:\n") //EXIBE EM ORDEM ALFABÉTICA CRESCENTE OU DECRESCENTE, DEPENDE DO QUE EU QUISER
	harmonia(personagens)
}

func harmonia(x []livros) {
	for i, v := range x {
		fmt.Println(i, "\tIdade:", v.Idade, "Personagem:", v.Personagem, "Livro:", v.Nome, "\n")  
  //EXIBE O INDICE (ONDE ESTÃO ARMAZENADOS NA SLICE) E TUDO O QUE EU PEDI DEPOIS
  
		for _, n := range v.Nomeperso {
			fmt.Println("\t", n) //EXIBE O NOME DOS PERSONAGENS 
		}
		fmt.Println("\n")
	}
}


///////////////////////////////////////////////////////////////////
                          MEU EXEMPLO 2 


package main

import (
	"fmt"
	"sort"
)

type comida struct {
	Nome   string
	Nota   int
	Motivo []string
}

type porNota []comida

func (p porNota) Len() int           { return len(p) }
func (p porNota) Less(i, j int) bool { return p[i].Nota < p[j].Nota }
func (p porNota) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type porNome []comida

func (p porNome) Len() int           { return len(p) }
func (p porNome) Less(i, j int) bool { return p[i].Nome < p[j].Nome }
func (p porNome) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {

	c1 := comida{
		Nome: "Strogonof",
		Nota: 9,
		Motivo: []string{
			"Muito bom",
			"Pena que tem leite",
			"O de frango é melhor",
		},
	}
	c2 := comida{
		Nome: "Pizza",
		Nota: 10,
		Motivo: []string{
			"Sem defeitos",
			"A melhor é mexicana",
			"Todo mundo gosta",
		},
	}
	c3 := comida{
		Nome: "Quiabo",
		Nota: 0,
		Motivo: []string{
			"Muito ruim",
			"Quem gosta não tem amor a vida",
			"Nota dó de tanta tristeza",
		},
	}
	comidas := []comida{c1, c2, c3}
	fmt.Println(comidas)

	for _, v := range comidas {
		sort.Strings(v.Motivo)
	}

	sort.Sort(porNota(comidas))
	fmt.Println("\nOrdena pela nota:")
	organizacao(comidas)

	sort.Sort(porNome(comidas))
	fmt.Println("\nOrdena pelo nome:")
	organizacao(comidas)
}

func organizacao(x []comida) {
	for i, v := range x {
		fmt.Println(i, "\nNome:", v.Nome, "\nNota:", v.Nota, "\nMotivo da nota:", v.Motivo)
		for _, n := range v.Motivo {
			fmt.Println("\t", n)
		}
	}
}


//tela

//[{Strogonof 9 [Muito bom Pena que tem leite O de frango é melhor]} {Pizza 10 [Sem defeitos A melhor é mexicana Todo mundo gosta]} {Quiabo 0 [Muito ruim Quem gosta não tem amor a vida Nota dó de tanta tristeza]}]

//Ordena pela nota:
//0 
//Nome: Quiabo 
//Nota: 0 
//Motivo da nota: [Muito ruim Nota dó de tanta tristeza Quem gosta não tem amor a vida]
//	 Muito ruim
//	 Nota dó de tanta tristeza
//	 Quem gosta não tem amor a vida
//1 
//Nome: Strogonof 
//Nota: 9 
//Motivo da nota: [Muito bom O de frango é melhor Pena que tem leite]
//	 Muito bom
//	 O de frango é melhor
//	 Pena que tem leite
//2 
//Nome: Pizza 
//Nota: 10 
//Motivo da nota: [A melhor é mexicana Sem defeitos Todo mundo gosta]
//	 A melhor é mexicana
//	 Sem defeitos
//	 Todo mundo gosta

//Ordena pelo nome:
//0 
//Nome: Pizza 
//Nota: 10 
//Motivo da nota: [A melhor é mexicana Sem defeitos Todo mundo gosta]
//	 A melhor é mexicana
//	 Sem defeitos
//	 Todo mundo gosta
//1 
//Nome: Quiabo 
//Nota: 0 
//Motivo da nota: [Muito ruim Nota dó de tanta tristeza Quem gosta não tem amor a vida]
//	 Muito ruim
//	 Nota dó de tanta tristeza
//	 Quem gosta não tem amor a vida
//2 
//Nome: Strogonof 
//Nota: 9 
//Motivo da nota: [Muito bom O de frango é melhor Pena que tem leite]
//	 Muito bom
//	 O de frango é melhor
//	 Pena que tem leite



///////////////////////////////////////////////////////////////////

                         EXEMPLO DA AULA
package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type porIdade []user

func (p porIdade) Len() int           { return len(p) }    //retorna os elementos do slice
func (p porIdade) Less(i, j int) bool { return p[i].Age < p[j].Age } //retorna se um é mais novo(menor) que o outro e vice-versa
func (p porIdade) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }    
//troca a ordem dos elementos para que fiquem de acordo com o Less

type sobreNome []user

func (p sobreNome) Len() int           { return len(p) }  //retorna todos os elementos do slice
func (p sobreNome) Less(i, j int) bool { return p[i].Last < p[j].Last } //retorna se um é mais novo(menor) que o outro e vice-versa
func (p sobreNome) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }     
//troca a ordem dos elementos para que fiquem de acordo com o Less

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	for _, v := range users {
		sort.Strings(v.Sayings) //isso coloca em ordem alfabética as falas
	}

	sort.Sort(porIdade(users))
	fmt.Println("Ordena por idade:\n") //exibe a ordenação por idade
	harmoniosa(users)

	sort.Sort(sobreNome(users))
	fmt.Println("Ordena por sobrenome:\n") //exibe a ordenação por ordem alfabética de sobrenome
	harmoniosa(users)
}


//cria uma função para não precisar repetir isso todas as vezes que for imprimir a idade e sobrenome em ordem
func harmoniosa(x []user) {
	for i, v := range x {
		fmt.Println(i, "\tIdade:", v.Age, "\tNome completo:", v.First, v.Last, "\n")  //imprime a idade, nome e sobrenome
		for _, c := range v.Sayings {
			fmt.Println("\t", c) //imprime as citações
		}
		fmt.Println("\n")  //só pula uma linha

	}
}

//tela
//Ordena por idade:

//0 	Idade: 27 	Nome completo: Miss Moneypenny 

//	 I would really prefer to be a secret agent myself.
//	 James, it is soo good to see you
//	 Would you like me to take care of that for you, James?


//1 	Idade: 32 	Nome completo: James Bond 

//	 In his majesty's royal service
//	 Shaken, not stirred
//	 Youth is no guarantee of innovation


//2 	Idade: 54 	Nome completo: M Hmmmm 

//	 Can someone please tell me where James Bond is?
//	 Dear God, what has James done now?
//	 Oh, James. You didn't.


//Ordena por sobrenome:

//0 	Idade: 32 	Nome completo: James Bond 

//	 In his majesty's royal service
//	 Shaken, not stirred
//	 Youth is no guarantee of innovation


//1 	Idade: 54 	Nome completo: M Hmmmm 

//	 Can someone please tell me where James Bond is?
//	 Dear God, what has James done now?
//	 Oh, James. You didn't.


//2 	Idade: 27 	Nome completo: Miss Moneypenny 

//	 I would really prefer to be a secret agent myself.
//	 James, it is soo good to see you
//	 Would you like me to take care of that for you, James?
