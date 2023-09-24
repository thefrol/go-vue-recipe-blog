package credentials

func ValidateCredentials(login string, pass string) bool {
	if login == "thefrol" && pass == "mypass" {
		return true
	}
	return false
}

const secretToken = "my.secret.token.101"

func CheckToken(token string) bool {
	if token == secretToken {
		return true
	}
	return false
}

func MakeToken() string {
	return secretToken
}
