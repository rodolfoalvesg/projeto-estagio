package checkemail

// Função verifica se existe email repetido
func EmailRepetido(email ListEmails, emailTest string) int {
	soma := 0
	for _, value := range email {

		if value == emailTest {
			soma++
		}
	}
	return soma
}
