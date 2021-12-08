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
		"",
		"teste4@teste.com",
		"teste5@teste.com",
	}
)

// Função de conversão do tipo string vindo da matriz, para inteiro
func Convert(qtd, price string) int {
	intQtd, _ := strconv.ParseInt(qtd, 0, 0)
	intPrice, _ := strconv.ParseInt(price, 0, 0)

	sub := int(intQtd * intPrice)

	return sub
}

//  Método para calcular a soma do Valor total e a divisão por usuários
func (d Dicionario) CalcList(valor ShopList, email ListEmails) Dicionario {

	teste := false

	// Veririca se existe email repetido
	for _, value := range email {
		if EmailRepetido(email, value) > 1 {
			teste = true
		}
	}

	// valida a verificação anterior
	if teste == false {
		soma := 0
		for _, value := range valor {
			soma += Convert(value[1], value[2])
		}

		resto := (soma % len(email)) - 1
		divisao := soma / len(email)

		for i, value := range email {

			if resto < i {
				d[value] = strconv.Itoa(divisao)
			} else {
				d[value] = strconv.Itoa(divisao + 1)
			}

		}

		return d
	} else {
		fmt.Printf("existe um email repetido \n")
		return d
	}

}

// Função verifica se existe email repetido
func EmailRepetido(email ListEmails, emailTest string) int {
	soma := 0
	for _, value := range email {

		if value == emailTest {
			soma++
		}
	}
	return soma
}

/* Verificação se existe Posições na lista de email vazias*/
func VerificaPosicaoVazia() bool {

	vazio := false
	for _, value := range EmailsList {
		if value == "" {
			vazio = true
		}
	}
	return vazio
}

// Funcão principal com a chamada de teste
func main() {

	if (len(EmailsList) != 0) && (VerificaPosicaoVazia() != true) {
		lislando := dicionario.CalcList(MyList, EmailsList)
		fmt.Printf("Dicionário: '%v'\n", lislando)

	} else {
		fmt.Printf("\nA lista de e-mail está vazia ou existe uma posição com valor vazio\n")
	}

}
