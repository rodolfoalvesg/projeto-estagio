package validates

type Emails []string

// validação de email repetido ou campo vazio
func (e Emails) Validate() bool {

	emailRepeated := make(map[string]struct{}, len(e))

	for _, entry := range e {
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
