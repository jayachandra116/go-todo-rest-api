package utils

import (
	"regexp"
)

// whether an email address is valid or not
func IsValidEmail(email string) bool {
	emailRegexp, _ := regexp.Compile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegexp.MatchString(email)
}

// min length = 4
func IsValidPassword(password string) bool {
	return len(password) >= 4
}
