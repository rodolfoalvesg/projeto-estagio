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
		Price:  250,
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
	got := ListCalculator(items, emails)
	want := Dictionary{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}
