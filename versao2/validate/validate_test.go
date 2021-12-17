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

	testValidate := map[string]struct {
		emails EmailsList
		want   bool
	}{
		"Correto":     {EmailsList{emailsA}, false},
		"Email duplo": {EmailsList{emailsB}, true},
		"Email vazio": {EmailsList{emailsC}, true},
	}

	for name, tt := range testValidate {
		got := tt.emails.ValidateEntries()
		if got != tt.want {
			t.Errorf("%s:  got %v, want %v", name, got, tt.want)
		}
	}

}
