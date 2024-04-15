package session

import (
	"crypto/rsa"
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GetPublicSessionKey() *rsa.PublicKey {
	publicKeyPath := "../../test/jwtRS256.key.pub.pem"
	if env, ok := os.LookupEnv("SESSION_PUB_KEY"); ok {
		publicKeyPath = env
	} else {
		log.Println("Using test keys for session!!!")
	}
	publicKey, err := os.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatalf("could not load session key: %s", err)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		log.Fatalf("could not load session key: %s", err)
	}
	return key
}

type CustomClaims struct {
	UserId string
	jwt.RegisteredClaims
}

func GetClaims(signedToken string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return GetPublicSessionKey(), nil
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return token.Claims.(*CustomClaims), nil
}

func CreateSignedToken(userId string) (string, error) {
	privateKeyPath := "../../test/jwtRS256.key"
	if env, ok := os.LookupEnv("SESSION_PUB_KEY"); ok {
		privateKeyPath = env
	} else {
		log.Println("Using test keys for session!!!")
	}
	privateKey, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return "", err
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, CustomClaims{
		UserId: userId,
	})
	tokenString, err := jwtToken.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
