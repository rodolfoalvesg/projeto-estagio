package validate

import (
	"testing"
)

var emailsA = []string{
	"teste1@teste.com",
	"teste2@teste.com",
	"teste3@teste.com",
	"teste4@teste.com",
}

var emailsB = []string{
	"teste1@teste.com",
	"teste2@teste.com",
	"teste3@teste.com",
	"teste3@teste.com",
}

var emailsC = []string{
	"teste1@teste.com",
	"teste2@teste.com",
	"",
	"teste3@teste.com",
}

func TestValidateEmail(t *testing.T) {

	validateEmails := func(t *testing.T, got, want bool) {
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("E-mail Ok", func(t *testing.T) {
		got := ValidateEntries(emailsA)
		validateEmails(t, got, false)
	})

	t.Run("E-mail duplicado", func(t *testing.T) {
		got := ValidateEntries(emailsB)
		validateEmails(t, got, true)
	})

	t.Run("E-mail vazio", func(t *testing.T) {
		got := ValidateEntries(emailsC)
		validateEmails(t, got, true)
	})

}
