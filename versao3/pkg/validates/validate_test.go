package validates

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
		emails Emails
		want   bool
	}{
		"Correto":     {emailsA, false},
		"Email duplo": {emailsB, true},
		"Email vazio": {emailsC, true},
	}

	for name, tt := range testValidate {
		got := tt.emails.Validate()
		if got != tt.want {
			t.Errorf("%s:  got %v, want %v", name, got, tt.want)
		}
	}

}
