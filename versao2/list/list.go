package list

type ListItens struct {
	Item, Spec    string
	Amount, Price uint
}

var BaseItens = []ListItens{
	{
		Item:   "Cenoura",
		Amount: 2,
		Price:  500,
		Spec:   "KG",
	},
	{
		Item:   "Batata",
		Amount: 2,
		Price:  250,
		Spec:   "KG",
	},
	{
		Item:   "Carne",
		Amount: 2,
		Price:  1500,
		Spec:   "KG",
	},
	{
		Item:   "Óleo",
		Amount: 1,
		Price:  253,
		Spec:   "UNI",
	},
	{
		Item:   "Macarrão",
		Amount: 2,
		Price:  130,
		Spec:   "PCT",
	},
	{
		Item:   "Biscoito",
		Amount: 5,
		Price:  250,
		Spec:   "PCT",
	},
}

var BaseEmails = []string{
	"teste1@teste.com",
	"teste2@teste.com",
	"teste3@teste.com",
	"teste4@teste.com",
}
