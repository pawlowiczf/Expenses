package token

import "time"

type TokenMaker interface {
	CreateToken(userID int64, email string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}