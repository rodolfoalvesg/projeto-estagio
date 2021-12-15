package validate

import (
	"testing"
)

func TestValidateEmail(t *testing.T) {
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

	testeValidate := map[string]struct {
		emails Emails
		want   bool
	}{
		"Correto":     {EmailsList{emailsA}, false},
		"Email duplo": {EmailsList{emailsB}, true},
		"Email vazio": {EmailsList{emailsC}, true},
	}

	t.Run("E-mail Ok", func(t *testing.T) {

		for name, tt := range testeValidate {
			got := tt.emails.ValidateEntries()
			if got != tt.want {
				t.Errorf("%s:  got %v, want %v", name, got, tt.want)
			}
		}

	})

}
