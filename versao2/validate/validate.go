package validate

type Emails interface {
	ValidateEntries() bool
}

type EmailsList struct {
	Emails []string
}

// validação de email repetido ou campo vazio
func (e EmailsList) ValidateEntries() bool {

	emailRepeated := map[string]string{}

	for _, entry := range e.Emails {
		if _, ok := emailRepeated[entry]; !ok {
			emailRepeated[entry] = ""
		} else {
			return true
		}

		if entry == "" {
			return true
		}
	}
	return false
}
