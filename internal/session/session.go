package session

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GetSessionKey() []byte {
	if envKey, ok := os.LookupEnv("SESSION_KEY"); ok {
		return []byte(envKey)
	}
	log.Println("Using test session key!")
	return []byte("Session Secret Key")
}

type CustomClaims struct {
	UserId string
	jwt.RegisteredClaims
}

func GetClaims(signedToken string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return GetSessionKey(), nil
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return token.Claims.(*CustomClaims), nil
}
