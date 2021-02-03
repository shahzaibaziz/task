package JWT

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(Email string) (string, error) {
	JWTSigningKey := []byte("thisIstheSuperSecretKey")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = Email

	tokenString, err := token.SignedString(JWTSigningKey)

	if err != nil {
		panic(err.Error())
		fmt.Errorf("Something has gone: %s", err.Error())
		return " ", err
	}
	return tokenString, nil
}