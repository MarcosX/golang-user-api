package db

import (
	"errors"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func getSessionKey() []byte {
	if key, ok := os.LookupEnv("SESSION_KEY"); ok {
		return []byte(key)
	}
	log.Println("Using test session key")
	return []byte("Session Secret Key")
}

type Session struct {
	SignedToken string
}

func NewSession(userId string) *Session {
	panic("Not implemented")
}

type CustomClaims struct {
	UserId string
	jwt.RegisteredClaims
}

func (s *Session) GetClaims() (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(s.SignedToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return getSessionKey(), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims, nil
	}
	return nil, errors.New("could not parse claims")
}

type SessionRepository interface {
	GetSession(id string) (*Session, error)
}

type realSessionRepository struct {
}

func (u *realSessionRepository) GetSession(sessionId string) (*Session, error) {
	panic("Not implemented")
}

func NewSessionRepository() SessionRepository {
	return &realSessionRepository{}
}
