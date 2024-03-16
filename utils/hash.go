package utils

import "golang.org/x/crypto/bcrypt"

func GetHashedPassword(plainTextPassword string) (string, error) {
	byteHashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), 14)
	return string(byteHashedPassword), err
}

func CheckPasswordHash(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
