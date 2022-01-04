package main

import (
	"fmt"

	"github.com/rodolfoalvesg/projeto-estagio/versao2/calculator"
	"github.com/rodolfoalvesg/projeto-estagio/versao2/list"
	"github.com/rodolfoalvesg/projeto-estagio/versao2/validate"
)

func main() {
	if (len(list.BaseEmails) != 0) && validate.Emails.Validate(list.BaseEmails) != true {
		// Recebe o calculo do total de cada e o resto
		valuePerPerson, remaining := calculator.ListCalculator(list.BaseItens, list.BaseEmails)

		// to Most cria uma instancia distribuindo os valores e o resto para maioria
		toMost := calculator.RemainingForMost{ValuePerPerson: valuePerPerson, Remaining: remaining, Email: list.BaseEmails}

		// to One cria uma instancia distribuindo os valores e o resto para um
		toOne := calculator.RemainingForOne{ValuePerPerson: valuePerPerson, Remaining: remaining, Email: list.BaseEmails}

		fmt.Println("--------- RESTO A MAIORIA --------")
		for index, person := range toMost.AddedValuesToUser() {
			fmt.Printf("E-mail %s: %d centavos\n", index, person)
		}

		fmt.Println("\n\n--------- RESTO A UM --------")
		for index, person := range toOne.AddedValuesToUser() {
			fmt.Printf("E-mail %s: %d centavos\n", index, person)
		}

	} else {

		fmt.Println("Existe um email vazio ou a sua lista de emails está vazia")

	}

}
