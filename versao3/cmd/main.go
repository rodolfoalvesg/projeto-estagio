package main

import (
	"fmt"

	"github.com/rodolfoalvesg/projeto-estagio/versao3/accounts_calculator"
	"github.com/rodolfoalvesg/projeto-estagio/versao3/lists"
	"github.com/rodolfoalvesg/projeto-estagio/versao3/validates"
)

func main() {
	if (len(lists.BaseEmails) != 0) && validates.Emails.Validate(lists.BaseEmails) != true {
		fmt.Println("--------- MAPA DE VALORES --------")
		listSales := accounts_calculator.ListCalculator(lists.BaseItens, lists.BaseEmails) // Passando as listas para as funções de calculo

		for index, person := range listSales {
			fmt.Printf("E-mail %s: %d centavos\n", index, person)
		}
	} else {
		fmt.Println("Existe um email vazio ou a sua lista de emails está vazia")
	}

}
