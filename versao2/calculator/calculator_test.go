package calculator

import (
	"modulo/list"
	"reflect"
	"testing"
)

var items = []list.ListItens{
	{
		Item:   "Arroz",
		Amount: 4,
		Price:  500,
		Spec:   "kg",
	},
	{
		Item:   "Macarrão",
		Amount: 2,
		Price:  251,
		Spec:   "pc",
	},
	{
		Item:   "Tomate",
		Amount: 5,
		Price:  100,
		Spec:   "un",
	},
}

var emails = []string{
	"teste1@teste.com",
	"teste2@teste.com",
	"teste3@teste.com",
}

func TestListCalculator(t *testing.T) {
	valuePerPerson := 1000
	remaining := 1

	t.Run("atribui", func(t *testing.T) {
		got := addedValuesUser(valuePerPerson, remaining, emails)
		want := Dictionary{
			"teste1@teste.com": 1001,
			"teste2@teste.com": 1000,
			"teste3@teste.com": 1000,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("calcular", func(t *testing.T) {
		got := ListCalculator(items, emails)
		want := Dictionary{
			"teste1@teste.com": 1001,
			"teste2@teste.com": 1001,
			"teste3@teste.com": 1000,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
