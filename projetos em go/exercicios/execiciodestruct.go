package main

import "fmt"

type marmita struct {
	base      string
	mistura   string
	bebida    string
	sobremesa bool
}

type acompanha struct {
	marmita
	buscar bool
	valor  float64
}

func main() {
	marmita1 := acompanha{
		marmita: marmita{
			base:      "arroz e feijão,",
			mistura:   "frango e batata,",
			bebida:    "suco de laranja,",
			sobremesa: true,
		},
		buscar: true,
		valor:  20.00,
	}
	marmita2 := acompanha{
		marmita: marmita{
			base:      "macarrão,",
			mistura:   "bife a milanesa,",
			bebida:    "coca-cola,",
			sobremesa: true,
		},
		buscar: false,
		valor:  22.00,
	}
	marmita3 := acompanha{
		marmita: marmita{
			base:      "arroz e feijoada,",
			mistura:   "torresmo e farofa,",
			sobremesa: false,
		},
		buscar: true,
		valor:  33.75,
	}
	fmt.Println(marmita1.mistura)  //assim eu mostro só a parte que digitei ali, so caso só mostra a mistura
	fmt.Println("\n", marmita2)
	fmt.Println("\n", marmita3)
}
//tela 
//frango e batata,
//{{macarrão, bife a milanesa, coca-cola, true} false 22}
//{{arroz e feijoada, torresmo e farofa,  false} true 33.75}
