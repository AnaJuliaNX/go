        		ALTERANDO DE GO PARA JSON

package main

import (
	"encoding/json"
	"fmt"
)

type pessoa struct {
	Nome      string
	Sobrenome string   //as letras iniciais precisam ser maisculuas para que possam ser visiveis e importadas
	Vivo      bool
	Idade     int
}

func main() {

	sam := pessoa{
		Nome:      "Sam",
		Sobrenome: "Cortland",
		Vivo:      false,
		Idade:     21,
	}
	aelin := pessoa{"Aelin", "Não sei", true, 19}

	s, err := json.Marshal(sam)
	if err != nil {
		fmt.Println(err)
	}

	a, err := json.Marshal(aelin)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(s))
	fmt.Println(string(a))
}

//tela
//{"Nome":"Sam","Sobrenome":"Cortland","Vivo":false,"Idade":21}
//{"Nome":"Aelin","Sobrenome":"Não sei","Vivo":true,"Idade":19}


/////////////////////////////////////////////////////////////////

                  ALTERANDO DE JSON PARA GO

package main

import (
	"encoding/json"
	"fmt"
)

type info struct {
	Nome      string `json:"Nome"`
	Sobrenome string `json:"Sobrenome"`
	Vivo      bool   `json:"Vivo"`    //isso são tags, servem para dizer na maioria das vezes que isso em Json tem um nome e em go
	Idade     int    `json:"Idade"`   //pode ter outro
}

func main() {

	sb := []byte(`{"Nome":"Sam","Sobrenome":"Cortland","Vivo":false,"Idade":21}`) //isso é o codigo em json

	var sam info
	err := json.Unmarshal(sb, &sam)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(sam)
	fmt.Println(sam.Sobrenome)
}
//tela
//{Sam Cortland false 21}
//Cortland



/////////////////////////////////////////////////////////////////

                 	 ENCODER

package main

import (
	"encoding/json"
	"os"
)
type pessoa struct {
	Nome      string
	Sobrenome string
	Vivo      bool
	Idade     int
}

func main() {

	sam := pessoa{
		Nome:      "Sam",
		Sobrenome: "Cortland",
		Vivo:      false,
		Idade:     21,
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(sam)
}
//tela
//{"Nome":"Sam","Sobrenome":"Cortland","Vivo":false,"Idade":21}

