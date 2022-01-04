package calculator

import (
	"github.com/rodolfoalvesg/projeto-estagio/versao2/list"
)

type totalPerClient map[string]uint

var clienteAndValues = totalPerClient{}

type ValueDistribuitor interface {
	AddedValuesToUser() map[string]uint
}

type RemainingForMost struct {
	ValuePerPerson int
	Remaining      int
	Email          []string
}

type RemainingForOne struct {
	ValuePerPerson int
	Remaining      int
	Email          []string
}

// Distribui os valores e emails num dicionário e o resto para MUITOS
func (r RemainingForMost) AddedValuesToUser() totalPerClient {
	for i, value := range r.Email {
		clienteAndValues[value] = uint(r.ValuePerPerson)
		if r.Remaining > i {
			clienteAndValues[value] += 1
		}
	}

	return clienteAndValues
}

// Distribui os valores e emails num dicionario e o resto para UM
func (r RemainingForOne) AddedValuesToUser() totalPerClient {
	for i, value := range r.Email {
		clienteAndValues[value] = uint(r.ValuePerPerson)
		if i == len(r.Email)-1 {
			clienteAndValues[value] = uint(r.ValuePerPerson) + uint(r.Remaining)
		}
	}

	return clienteAndValues
}

// Realiza o subcalculo da QTD x VALOR em seguida chama a função AddedValuesUser() para distribuir Os valores num dicionario
func ListCalculator(items []list.ListItens, emails []string) (int, int) {
	sum := 0
	for _, value := range items {
		sum += (int(value.Price) * int(value.Amount))
	}
	valuePerson := sum / len(emails)
	remaining := sum % len(emails)

	return valuePerson, remaining
}
