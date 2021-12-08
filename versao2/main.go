package main

import (
	"fmt"
	"modulo/calculator"
	"modulo/list"
	"modulo/validate"
)

func main() {
	items, emails := list.List() // recebendo a lista de ITENS e EMAILS

	if (len(emails) != 0) && validate.ValidateEntries(emails) != true {
		fmt.Println("--------- MAPA DE VALORES --------")
		calculator.ListCalculator(items, emails) // Passando as listas para as funções de calculo
	} else {
		fmt.Println("Existe um email vazio ou a sua lista de emails está vazia")
	}

}
