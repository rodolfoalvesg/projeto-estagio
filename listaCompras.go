package main

import (
	"fmt"
	"strconv"
)

//Tipos definidos
type (
	ShopList   [][]string
	ListEmails []string
	Dicionario map[string]string
)

// Dicionário vazio
var dicionario = Dicionario{}

// Listas
var (
	MyList = ShopList{
		{"Arroz", "4", "500", "wgt"},
		{"Macarrão", "2", "300", "pkg"},
		{"Feijão", "1", "806", "wgt"},
		{"Maça", "5", "100", "uni"},
	}

	EmailsList = ListEmails{
		"teste1@teste.com",
		"teste2@teste.com",
		"teste2@teste.com",
		"teste3@teste.com",
		"teste4@teste.com",
		"teste5@teste.com",
	}
)

// Função de conversão do tipo string vindo da matriz, para inteiro
func Convert(qtd, price string) int {
	intQtd, _ := strconv.ParseInt(qtd, 0, 0)
	intPrice, _ := strconv.ParseInt(price, 0, 0)

	sub := int(intQtd * intPrice)

	return sub // retorna total da multiplicacao da quantidade pelo valor do item
}

//  Método para calcular a soma do Valor total e a divisão por usuários
func (d Dicionario) CalcList(valor ShopList, email ListEmails) Dicionario {

	soma := 0 // inicializador de soma
	for _, value := range valor {
		soma += Convert(value[1], value[2])
	} // itera VALOR que é do tipo ShopLista e contem todos os itens da lista

	resto := (soma % len(email)) - 1
	divisao := soma / len(email)

	for i, value := range email {
		if EmailRepetido(email, value) > 1 {
			fmt.Printf("O email: '%s', está repetido", value)
			return d
		}

		if resto < i {
			d[value] = strconv.Itoa(divisao)
		} else {
			d[value] = strconv.Itoa(divisao + 1)
		}

	} // itera EMAIL que é do tipo ListEmails e contem todos os emails da lista

	return nil // receptor que contém a lista de usuários e os valores atribuídos
}

// Funcão principal com a chamada de teste
func main() {

	if len(EmailsList) != 0 {
		lislando := dicionario.CalcList(MyList, EmailsList)
		fmt.Printf("Dicionário: '%v'\n", lislando)

	} else {
		fmt.Printf("A lista de e-mail está vazia. Deve haver pelo menos um valor")
	}

}
