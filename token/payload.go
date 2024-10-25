package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Payload struct {
	ID     uuid.UUID `json:"id"`
	UserID int64     `json:"user_id"`
	Email  string    `json:"email"`

	jwt.RegisteredClaims
	// IssuedAt  time.Time `json:"issued_at"`
	// ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(userID int64, email string, duration time.Duration) (*Payload, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		ID:     uuid,
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(duration)},
			NotBefore: &jwt.NumericDate{Time: time.Now()},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			ID:        uuid.String(),
		},
		// IssuedAt: time.Now(),
		// ExpiredAt: time.Now().Add(duration),
	}, nil
}
