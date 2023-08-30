//ERRORS.NEW

package main

import (
	"errors"
	"log"
)

func main() {
	_, err := sqrt(-10) //chama uma função de raiz quadrada de -10
	if err != nil {
		log.Fatalln(err) //se tiver erro retorna um erro fatal/fatal err
	}
}

func sqrt(f float64) (float64, error) { //essa func toma um float e retorna um float e um erro
	if f < 0 { //se o f for menor que zero retorna o de baixo
		return 0, errors.New("Calculo impossivel: rais quadrada de número negativo")
	}
	return 42, nil //se não for menor que zero não retorna nada
}

////////////////////////////////////////////////////////////////////
//VAR ERRORS.NEW

package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrCalInval = errors.New("Calculo inválido: raiz quadrada de um número negativo")
//ao invés de usar só uma vez dentro do código coloco aqui fora e uso quantas vezes quiser

func main() {
	fmt.Printf("%T\n", ErrCalInval) //isso é só pra mostrar o tipo, confirmar que é um erro
	_, err := sqrt(-10)
	if err != nil {    //isso é a mesma coisa que o exemplo de cima
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrCalInval
	}
	return 42, nil
}

////////////////////////////////////////////////////////////////////
//FMT.ERROF

package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10.23)
	if err != nil {    //mesma coisa do primeiro, executa e retorna se tem erro
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("Calculo inválido novamente: raiz de número negativo: %v", f)
//mas nesse caso ele vai me dizer qual número que deu BO
	}
	return 42, nil
}

////////////////////////////////////////////////////////////////////
//VAR FMT.ERRORF

package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10.23)
	if err != nil { //mesma coisa do primeiro
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		ErrCalInval := fmt.Errorf("Calculo inválido novamente: raiz de número negativo: %v", f)
//mas nesse caso o meu ErrCalInval poderia ser uma váriavel antes de tudo e ainda eu usaria
//mais de uma vez
//mas é uma coisa um pouco especifica
		return 0, ErrCalInval
	}
	return 42, nil
}

////////////////////////////////////////////////////////////////////
//TYPE + METHOD = ERROR INTERFACE

package main

import (
	"fmt"
	"log"
)

type erromatrmatico struct { 
	num1 string
	num2 string
	err  error
}
//se não tivesse o método eu não poderia usar a struct com um tipo erro
func (n erromatrmatico) Error() string {  //temos um método e um erro  que retorna uma string
	return fmt.Sprintf("Ocorreu um erro de calculo: %v %v %v", n.num1, n.num2, n.err)
	//o0s valores do num1 e num2 estão colocados dentro da func sqrt
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil { //mesma função da primeira
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("\nCalculo invalido reduzido: raiz de número negativo: %v", f)
		//criei uma var que tem um erro novo onde vou passar o valor
		return 0, erromatrmatico{"3.14 N1", "7.77 N2", nme}
		//mas não vou simplesmente retornar o erro, vou retornar a struct com esses valores e o erro
	}
	return 42, nil
}
