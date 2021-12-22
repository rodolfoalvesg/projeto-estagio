package calculator

import (
	"reflect"
	"testing"

	"github.com/rodolfoalvesg/projeto-estagio/versao2/list"
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

func TestAddedValuesUser(t *testing.T) {

	testCalculator := map[string]struct {
		valuePerPerson int
		remaining      int
		want           TotalPerClient
	}{
		"Teste Padrão": {
			valuePerPerson: 2000,
			remaining:      0,
			want: TotalPerClient{
				"teste1@teste.com": 2000,
				"teste2@teste.com": 2000,
				"teste3@teste.com": 2000,
			},
		},
		"Teste com Resto": {
			valuePerPerson: 1500,
			remaining:      2,
			want: TotalPerClient{
				"teste1@teste.com": 1501,
				"teste2@teste.com": 1501,
				"teste3@teste.com": 1500,
			},
		},
	}

	for name, tt := range testCalculator {
		got := addedValuesToUser(tt.valuePerPerson, tt.remaining, emails)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: ,got %v, want %v", name, got, tt.want)
		}
	}

}

func TestLisCalculator(t *testing.T) {
	got := ListCalculator(items, emails)
	want := TotalPerClient{
		"teste1@teste.com": 1001,
		"teste2@teste.com": 1001,
		"teste3@teste.com": 1000,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
