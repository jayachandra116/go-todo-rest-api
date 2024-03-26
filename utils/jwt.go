package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecretkey"

// Creates a new token from user's email and userId
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":   userId,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
			"email": email,
		})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// return users's userId,nil if the token is valid or nil,err
func VerifyToken(tokenString string) (int64, error) {
	fmt.Println("Token: ", tokenString)
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected SignatureMethod")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, errors.New("could not parse token")
	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid Token")
	}
	// Extract the data from the token
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	userId := int64(claims["sub"].(float64))
	return userId, nil
}
