                     //PRIMEIRO TESTE

package main

import "fmt"

func main() {

	testeum(umargumento)

}

func umargumento() {
	fmt.Println("\nEsse é o meu argumento")
	fmt.Println("Ele é escrito antes e exibido depois")
	fmt.Println("Isso tem sentido se analisar o código")
}

func testeum(primeiro func()) {
	fmt.Println("Essa é a minha função que recebe o argumento")
	fmt.Println("Ele é escrito depois e exibido antes")
	primeiro()

}
//tela
//Essa é minha função que recebe o argumento
//Ele é escrito depois e exibido antes

//Esse é o meu argumento
//Ele é escrito antes e exibe depois
//Isso tem sentido se analisar o código


//////////////////////////////////////////////////////////

                     //SEGUNDO TESTE

package main

import "fmt"

func main() {

	novoteste(novoargumento)

}

func novoargumento() {
	fmt.Println("\nNão aguento mais isso")
	fmt.Println("Mas preciso treinar pra tentar aprender")
	fmt.Println("Acho que ta dando certo")
}
func novoteste(novo func()) {
	fmt.Println("Meu braço já ta doendo")
	fmt.Println("Doendo de tanto digitar olha só")
	novo()
}
//tela
//Meu braço já ta doendo
//Doendo de tanto digitar olha só

//Não aguento mais isso
//Mas preciso treinar pra tentar aprender
//Acho que ta dando certo


//////////////////////////////////////////////////////////

                    // TERCEIRO TESTE

package main

import "fmt"

func main() {

	terceitoteste(terceiroargumento)
}

func terceiroargumento() {
	fmt.Println("\nEstou escrevendo tudo antes")
	fmt.Println("Pra ser exibido tudo depois")
	fmt.Println("Não faz muito sentido pra quem vê de fora mas ok")
}

func terceitoteste(terceiro func()) {
	fmt.Println("Estou escrevendo tudo depois")
	fmt.Println("Pra ser exibido tudo antes")
	terceiro()
}
//tela
//estou escrevendo tudo depois
//Pra ser exibido tudo antes

//Estou escrevendo tudo antes
//Pra ser exibido tudo depois
//Não faz muito sentido pra quem vê de fora mas ok


//////////////////////////////////////////////////////////

                   //  ULTIMO TESTE

package main

import "fmt"

func main() {

	ultimoteste(ultimoargumento)

}

func ultimoargumento() {
	fmt.Println("\nUltimo porque meus dedos estão doendo")
	fmt.Println("Não faço mais ideia do que digitar")
	fmt.Println("É isso acabou, até mais")
}
func ultimoteste(ultimo func()) {
	fmt.Println("É o ultimo por hora")
	fmt.Println("Ultimo até eu precisar usar isso de novo")
	fmt.Println("Ou até que tenha esquecido")
	ultimo()
}
//tela
//É o ultimo por hora
//Ultimo até eu precisar usar isso de novo
//Ou até que tenha esquecido

//Ultimo porque meus dedos estão doendo
//Não faço mais ideia do que digitar
//É isso acabou, até mais
