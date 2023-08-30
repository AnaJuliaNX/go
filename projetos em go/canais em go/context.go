//BACKGROUND

package main

import (
	"context"
	"fmt"
)

func main() {

	ctx := context.Background()

	fmt.Println("context:\t", ctx)//conteudo do meu context
	fmt.Println("context err:\t", ctx.Err()) //vai mostrar se tem algum erro, normalmente isso seria depois do
//cancel nesse caso como não estáe ele não exibe nenhum erro
	fmt.Printf("context type:\t%T\n", ctx)//nesse o tipo dele é "emptyCtx"
	fmt.Println("------------")
}

//////////////////////////////////////////////////////////////////
//WITHCANCEL

package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err\t", ctx.Err())   //isso aqui é a mesma coisa de antes
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("-----------")

	ctx, _ = context.WithCancel(ctx)

	fmt.Println("context\t", ctx) //conteudo do meu context
	fmt.Println("context err:\t", ctx.Err()) //exibiria o erro se eu já tivesse criado o cancel
	fmt.Printf("context type:\t%T\n", ctx) //tipo dele, nesse o tipo é "cancelCtx"
	fmt.Println("-----------")
}

//////////////////////////////////////////////////////////////////
//FUNÇÃO CANCEL

package main

import (
	"context"
	"fmt"
)

func main() {

	ctx := context.Background()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err()) //isso é a mesma coisa que já fiz um pouco acima
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("--------")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err()) //aqui não vou ter um erro já que não tem nada
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t\t", cancel)       //valor do cancel
	fmt.Printf("cancel type:\t%T\n", cancel) //tipo dele, nesse o tipo é "CancelFunc"
	fmt.Println("---------")

	cancel() //chamei o cancel pra rodar

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err()) //aqui vou ter um "erro", no caso vai exibir que foi cancelado
	fmt.Printf("context type:\t%T\n", ctx)   //aqui é o tipo dele, no caso "cancelCtx"
	fmt.Println("--------")
}
