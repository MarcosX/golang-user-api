package session

import (
	"crypto/rsa"
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type sessionData struct {
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}

var session *sessionData

func SessionData() *sessionData {
	if session != nil {
		return session
	}

	session = &sessionData{}

	publicKeyPath := "../../test/jwtRS256.key.pub.pem"
	if env, ok := os.LookupEnv("SESSION_PUBLIC_KEY"); ok {
		publicKeyPath = env
	} else {
		log.Println("Using test session keys!!!")
	}
	publicKeyFile, err := os.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatalf("could not load session key: %s", err)
	}
	session.PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyFile)
	if err != nil {
		log.Fatalf("could not load session key: %s", err)
	}

	privateKeyPath := "../../test/jwtRS256.key"
	if env, ok := os.LookupEnv("SESSION_PRIVATE_KEY"); ok {
		privateKeyPath = env
	} else {
		log.Println("Using test session keys!!!")
	}
	privateKeyFile, err := os.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("could not load session key: %s", err)
	}
	session.PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyFile)
	if err != nil {
		log.Fatalf("could not load session key: %s", err)
	}

	return session
}

func (s *sessionData) CreateSignedToken(userEmail string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, CustomClaims{
		UserEmail: userEmail,
	})
	tokenString, err := jwtToken.SignedString(s.PrivateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func EnforceValidSession() echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(CustomClaims)
		},
		ContextKey:    "user",
		SigningMethod: "RS256",
		SigningKey:    SessionData().PublicKey,
	}
	return echojwt.WithConfig(config)
}

type CustomClaims struct {
	UserEmail string
	jwt.RegisteredClaims
}

func (s *sessionData) ReadClaims(signedToken string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.PublicKey, nil
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return token.Claims.(*CustomClaims), nil
}
