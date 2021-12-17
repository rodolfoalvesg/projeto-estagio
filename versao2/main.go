package main

import (
	"fmt"

	"github.com/rodolfoalvesg/projeto-estagio/versao2/calculator"
	"github.com/rodolfoalvesg/projeto-estagio/versao2/list"
	"github.com/rodolfoalvesg/projeto-estagio/versao2/validate"
)

func main() {
	items, emails := list.List() // recebendo a lista de ITENS e EMAILS

	if (len(emails) != 0) && validate.EmailsList.ValidateEntries(validate.EmailsList{emails}) != true {
		fmt.Println("--------- MAPA DE VALORES --------")
		listSales := calculator.ListCalculator(items, emails) // Passando as listas para as funções de calculo

		for index, person := range listSales {
			fmt.Printf("E-mail %s: %d centavos\n", index, person)
		}
	} else {
		fmt.Println("Existe um email vazio ou a sua lista de emails está vazia")
	}

}
