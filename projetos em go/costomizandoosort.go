package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome     string
	potencia int
	consumo  int
}

type ordenaPotencia []carro  //slice de carro

func (x ordenaPotencia) Len() int           { return len(x) } //retorna os elementos tem na "slice ordenaPotencia"
func (x ordenaPotencia) Less(i, j int) bool { return x[i].potencia < x[j].potencia } //retorna se um elemento vem antes do outro
func (x ordenaPotencia) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }  //troca a ordem dos elementos

type ordenaConsumo []carro //slice de carro

func (x ordenaConsumo) Len() int           { return len(x) }  //retorna os elementos tem na "slice ordenaConsumo"
func (x ordenaConsumo) Less(i, j int) bool { return x[i].consumo > x[j].consumo } //retorna se um elemento vem antes do outro
func (x ordenaConsumo) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }  //troca a ordem dos elementos

type OrdenaLucroPosto []carro //slice de carro

func (x OrdenaLucroPosto) Len() int           { return len(x) } //retorna os elementos tem na "slice OrdenaLucroPosto"
func (x OrdenaLucroPosto) Less(i, j int) bool { return x[i].consumo < x[j].consumo } //retorna se um elemento vem antes do outro
func (x OrdenaLucroPosto) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }  //troca a ordem dos elementos

func main() {

	carros := []carro{carro{"Ferrari", 719, 20},
		carro{"Mercedes", 300, 15},
		carro{"Audi", 252, 19},
	}

	fmt.Println("Começa:\n", carros)

	sort.Sort(ordenaPotencia(carros))
	fmt.Println("\nPotencia:\n", carros)

	sort.Sort(ordenaConsumo(carros))
	fmt.Println("\nConsumo:\n", carros)

	sort.Sort(OrdenaLucroPosto(carros))
	fmt.Println("\nLucro do posto:\n", carros)
}
//tela
//Começa:
// [{Ferrari 719 20} {Mercedes 300 15} {Audi 252 19}]

//Potencia:
// [{Audi 252 19} {Mercedes 300 15} {Ferrari 719 20}]

//Consumo:
// [{Ferrari 719 20} {Audi 252 19} {Mercedes 300 15}]

//Lucro do posto:
// [{Mercedes 300 15} {Audi 252 19} {Ferrari 719 20}]
