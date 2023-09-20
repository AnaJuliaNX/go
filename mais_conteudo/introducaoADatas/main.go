package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	// Horário local
	now := time.Now()
	//Formatando como usamos aqui no Brasil
	fmt.Println("Formato padrão do Go:", now)
	fmt.Println("\nFormato personalizado:", now.Format("02/01/2006, 03:04:05"))

	//Para exibir de forma isolada os dados
	fmt.Println("\nDia:", now.Day())
	fmt.Println("Mês:", now.Month())
	fmt.Println("Ano:", now.Year())
	fmt.Println("Hora:", now.Hour())
	fmt.Println("Minutos", now.Minute())
	fmt.Println("Segundos:", now.Second())
	fmt.Println("Dia, mês e ano:", now.Day(), now.Month(), now.Year())

	//Para exibir uma data e hora pré determinada
	determinado := time.Date(2005, 03, 03, 23, 50, 0, 0, time.Local)
	fmt.Println("\n", determinado)
	fmt.Println("Formato personalizado:", determinado.Format("02/01/2006, 15:04:05"))

	//Outra forma de criar a data manualmente, tenho uma string e passo com time.Time
	datasalva := "23/10/2006 11:20:10 -03"
	parsed, erro := time.Parse("02/01/2006 03:04:05 -07", datasalva)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("\nData selecionada:", parsed.Format("02/01/2006 03:04:05 -07"))

	//Para saber quando vair ser daqui determinado tempo:
	maisDias := parsed.Add(15 * 24 * time.Hour)
	fmt.Println("Avançando 15 dias:", maisDias.Format("02/01/2006 03:04:05 -07"))

	//Para saber quanto tempo antes:
	menosDias := parsed.Add(-13 * 24 * time.Hour)
	fmt.Println("Voltando 13 dias:", menosDias.Format("02/01/2006 03:04:05 -07"))

	//Comparando datas usando bool:
	fmt.Println("\nVerificando se as datas são iguais:", maisDias.Equal(menosDias))

	//ATENÇÃO: A data de fora é sempre comparada com a de dentro
	//Comparando se uma data veio antes da outra:
	fmt.Println("Verificando se veio antes:", menosDias.Before(maisDias))
	//Comparando se uma data é depois da outra
	fmt.Println("Verificando se veio depois:", menosDias.After(maisDias))

}
