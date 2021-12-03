package main

import "testing"

var listaProd = [][]string{
	{"Arroz", "4", "500", "wgt"},
	{"Macarrão", "2", "300", "pkg"},
}
var listaEmail = []string{
	"teste1@teste.com",
	"teste2@teste.com",
}

func TesAdd(t *testing.T) {
	t.Run("lista nova", func(t *testing.T) {
		dicionario := Dicionario{}
		resultado, err := dicionario.CalcList(listaProd, listaEmail)

		if err != nil {
			t.Fatalf("deveria ter encontrado palavra adicionada:", err)
		}

	})
}
