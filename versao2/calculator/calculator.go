package calculator

import (
	"modulo/list"
)

type Dictionary map[string]uint

var dictionary = Dictionary{}

// Distribui os valores e emails num dicionário
func addedValuesUser(v1, v2 int, email []string) Dictionary {
	for i, value := range email {
		if v2 <= i {
			dictionary[value] = uint(v1)
		} else {
			dictionary[value] = uint(v1 + 1)
		}
	}

	return dictionary

}

// Realiza o subcalculo da QTD x VALOR em seguida chama a função AddedValuesUser() para distribuir
// Os valores num dicionario
func ListCalculator(items []list.ListItens, emails []string) Dictionary {
	sum := 0
	for _, value := range items {
		sum += (int(value.Price) * int(value.Amount))
	}
	valuePerson := sum / len(emails)
	remnant := sum % len(emails)

	return addedValuesUser(valuePerson, remnant, emails)
}
