                  DE GO PARA JSON


package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string      //as letras precisam começar sendo maiusculas 
	Age   int
}

func main() {

	u1 := user{
		First: "James",   //as letras precisam começar sendo maiusculas 
		Age:   32,
	}
	u2 := user{
		First: "Moneypenny",    //as letras precisam começar sendo maiusculas 
		Age:   27,
	}
	u3 := user{
		First: "M",    //as letras precisam começar sendo maiusculas 
		Age:   54,      
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
	a, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(a))
}

//tela
//[{James 32} {Moneypenny 27} {M 54}]
//[{"First":"James","Age":32},{"First":"Moneypenny","Age":27},{"First":"M","Age":54}]


////////////////////////////////////////////////////////////
                  DE JSON PARA GO

package main

import (
	"encoding/json"
	"fmt"
)

type jsontogo []struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	var resultado jsontogo

	err := json.Unmarshal([]byte(s), &resultado)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(resultado)    //o resultado é slice of structs, ou seja, são todos os resultados, um slice com vários structs
	fmt.Println(resultado[1]) //o resultado é um struct da slice
	fmt.Println(resultado[1].Last)
        // esse é como se fosse: item -> struct -> structs -> slice ->convertido de json para go

}

//tela
//[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]
//ISSO ESTÁ EM JSON

//[{James Bond 32 [Shaken, not stirred Youth is no guarantee of innovation In his majesty's royal service]} {Miss Moneypenny 27 [James, it is soo good to see you Would you like me to take care of that for you, James? I would really prefer to be a secret agent myself.]} {M Hmmmm 54 [Oh, James. You didn't. Dear God, what has James done now? Can someone please tell me where James Bond is?]}]
//ISSO FOI DE JSON PARA GO, OU SEJA, UM SLICE COM VÁRIOS STRUCTS

//{Miss Moneypenny 27 [James, it is soo good to see you Would you like me to take care of that for you, James? I would really prefer to be a secret agent myself.]}
//ISSO É APENAS UM ITEM DO MEU SLICE OF STRUCT

//Moneypenny
//UM ITEM DE UM STRUCT QUE FAZ PARTE DE VÁRIOS STRUCTS QUE FAZ PARTE DE UMA SLICE


////////////////////////////////////////////////////////////

                NEWENCODE E ENCODE PARA STDOUT

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

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

	fmt.Println(users)

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(users)
}
//tela 
//primeiro mostrou em go
//depois mostrou em json

////////////////////////////////////////////////////////////

             ORDEM ALFABÉTICA E NUMÉRICA NAS SLICE

package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)	

	sort.Strings(xs)
	fmt.Println(xs)
}
//tela
//[5 8 2 43 17 987 14 12 21 1 4 2 3 93 13]
//[1 2 2 3 4 5 8 12 13 14 17 21 43 93 987]
//[random rainbow delights in torpedo summers under gallantry fragmented moons across magenta]
//[across delights fragmented gallantry in magenta moons rainbow random summers torpedo under]
