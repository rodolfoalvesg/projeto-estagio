package validate

type EmailsList struct {
	Emails []string
}

// validação de email repetido ou campo vazio
func (e EmailsList) ValidateEntries() bool {

	emailRepeated := make(map[string]struct{}, len(e.Emails))

	for _, entry := range e.Emails {
		if _, ok := emailRepeated[entry]; ok {
			return true
		}

		emailRepeated[entry] = struct{}{}

		if entry == "" {
			return true
		}
	}

	return false
}
