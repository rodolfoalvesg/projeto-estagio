package calculator

import (
	"github.com/rodolfoalvesg/projeto-estagio/versao2/list"
)

type TotalPerClient map[string]uint

var clienteAndValues = TotalPerClient{}

// Distribui os valores e emails num dicionário
func addedValuesUser(valuePerPerson, remaining int, email []string) TotalPerClient {
	for i, value := range email {
		if remaining <= i {
			clienteAndValues[value] = uint(valuePerPerson)
		} else {
			clienteAndValues[value] = uint(valuePerPerson + 1)
		}
	}

	return clienteAndValues

}

// Realiza o subcalculo da QTD x VALOR em seguida chama a função AddedValuesUser() para distribuir
// Os valores num dicionario
func ListCalculator(items []list.ListItens, emails []string) TotalPerClient {
	sum := 0
	for _, value := range items {
		sum += (int(value.Price) * int(value.Amount))
	}
	valuePerson := sum / len(emails)
	remaining := sum % len(emails)

	return addedValuesUser(valuePerson, remaining, emails)
}
