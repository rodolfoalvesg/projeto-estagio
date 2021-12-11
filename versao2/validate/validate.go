package validate

// Verificação de email repetido
func repeateEmail(email []string, value string) bool {
	sum := 0

	for _, entry := range email {
		if value == entry {
			sum++
			if sum > 1 {
				return true
			}
		}
	}
	return false
}

// Verificação de email vazio
func ValidateEntries(emails []string) bool {
	for _, entry := range emails {

		if repeateEmail(emails, entry) {
			return true
		}

		if entry == "" {
			return true
		}
	}

	return false
}
