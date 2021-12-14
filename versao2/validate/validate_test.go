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

	testeValidate := []struct {
		emails Emails
		want   bool
	}{
		{EmailsList{emailsA}, false},
		{EmailsList{emailsB}, true},
		{EmailsList{emailsC}, true},
	}

	t.Run("E-mail Ok", func(t *testing.T) {

		for _, tt := range testeValidate {
			got := tt.emails.ValidateEntries()
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		}

	})

}
