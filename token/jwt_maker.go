package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32 

func NewJWTMaker(secretKey string) (TokenMaker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size, must be at least %d characters", minSecretKeySize)
	}
	
	return &JWTMaker{
		secretKey: secretKey,
	}, nil
}

type JWTMaker struct {
	secretKey string
}

func (maker *JWTMaker) CreateToken(userID int64, email string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(userID, email, duration)
	if err != nil {
		return "", nil, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	if err != nil {
		return "", nil, err
	}

	return token, payload, nil
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	//
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Method.Alg())
		}
		return []byte(maker.secretKey), nil
	}

	payload := &Payload{}
	jwtToken, err := jwt.ParseWithClaims(token, payload, keyFunc)
	if err != nil || !jwtToken.Valid {
		return nil, err 
	}

	return payload, nil 
}
