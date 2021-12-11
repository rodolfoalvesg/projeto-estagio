package validate

import (
	"testing"
)

var emails = []string{
	"teste1@teste.com",
	"teste2@teste.com",
	"teste3@teste.com",
	"teste4@teste.com",
}

func TestValidateEmail(t *testing.T) {

	t.Run("E-mail Ok", func(t *testing.T) {
		got := ValidateEntries(emails)
		want := false

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
