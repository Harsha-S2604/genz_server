package validations

import (
	"regexp"
	"unicode"
)

func ValidateUserData(email string, password string) (bool, string) {
	if email == "" || password == "" {
		return false, "one or more field(s) is missing"
	}
	var reEmail = regexp.MustCompile(`^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$`)
	if !(reEmail.MatchString(email)) {
		return false, "Please provide a valid email"
	}

	var (
		upp, low, num, sym bool
		tot                uint8
	)
 
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false, "Password should be minimum of 6 characters, maximum of 20 characters and must contain at least one uppercase, one lowercase letter, one number and one special character."
		}
	}
 
	if !upp || !low || !num || !sym || tot < 6 || tot > 20 {
		return false, "Password should be minimum of 6 characters, maximum of 20 characters and must contain at least one uppercase, one lowercase letter, one number and one special character."
	} 
	return true, ""
}