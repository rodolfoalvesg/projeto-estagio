package calculator

import (
	"github.com/rodolfoalvesg/projeto-estagio/versao2/list"
)

type TotalPerClient map[string]uint

var clienteAndValues = TotalPerClient{}

// Distribui os valores e emails num dicionário
func addedValuesToUser(valuePerPerson, remaining int, email []string) TotalPerClient {
	for i, value := range email {
		clienteAndValues[value] = uint(valuePerPerson)
		if remaining > i {
			clienteAndValues[value] += 1
		}
	}

	// for i := 0; i < remaining; i++ {
	// 	clienteAndValues[email[i]] += 1
	// }

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

	return addedValuesToUser(valuePerson, remaining, emails)
}
