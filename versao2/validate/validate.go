package validate

type Emails interface {
	ValidateEntries() bool
}

type EmailsList struct {
	emails []string
}

// validação de email repetido ou campo vazio
func (e EmailsList) ValidateEntries() bool {

	teste := map[string]string{}

	for _, entry := range e.emails {
		if _, ok := teste[entry]; !ok {
			teste[entry] = ""
		} else {
			return true
		}

		if entry == "" {
			return true
		}
	}
	return false
}
