package main

import (
	"fmt"

	"github.com/rodolfoalvesg/projeto-estagio/versao2/calculator"
	"github.com/rodolfoalvesg/projeto-estagio/versao2/list"
	"github.com/rodolfoalvesg/projeto-estagio/versao2/validate"
)

func main() {
	if (len(list.BaseEmails) != 0) && validate.Emails.Validate(list.BaseEmails) != true {
		fmt.Println("--------- MAPA DE VALORES --------")
		listSales := calculator.ListCalculator(list.BaseItens, list.BaseEmails) // Passando as listas para as funções de calculo

		for index, person := range listSales {
			fmt.Printf("E-mail %s: %d centavos\n", index, person)
		}
	} else {
		fmt.Println("Existe um email vazio ou a sua lista de emails está vazia")
	}

}
