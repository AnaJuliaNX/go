package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	senha := "20julho1980"
	senhaerrada := "20julho19990"

	sb, err := bcrypt.GerateFromPassword([]byte(senha), 10) //isso faz com que uma senha hash sejá feita
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb)) //esse aqui onde pede pra gerar o hash

	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha)))
	//isso faz com que sejam comparadas a hash e a senha

	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senhaerrada)))
	//isso faz com que sejam comparadas a hash e a senha errada

}

//tela de exmplo
//$2a#10IdiioozoiHUIFjio208576IOHFOfj8756ndidbd83653EOSOufhfbodh  //hash da senha

//<nil>  //não retorna nada pq ta certo, retornaria um erro se estivesse errrada

//"hasedPassword is not the hash of the given password"  //mensagem de erro
//essa menagem significa que a hash gerada anteriormente não é a mesma hash da senha digitada (a senha errada no caso)
