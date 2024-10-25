package token

import (
	"expenses/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	userID := util.RandomInt(1, 100)
	email := util.RandomEmail()
	duration := time.Minute

	token, payloadA, err := maker.CreateToken(userID, email, duration)
	require.NoError(t, err)
	require.NotEmpty(t, payloadA)
	require.NotEmpty(t, token)

	payloadB, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payloadB)
	require.Equal(t, payloadA.UserID, payloadB.UserID)
	require.Equal(t, payloadA.Email, payloadB.Email)
	require.Equal(t, payloadA.ID, payloadB.ID)

	require.WithinDuration(t, time.Now(),  payloadB.IssuedAt.Time, time.Second)
	require.WithinDuration(t, time.Now(), payloadB.ExpiresAt.Time, duration)
	require.WithinDuration(t, payloadA.IssuedAt.Time.Add(duration), payloadB.ExpiresAt.Time, time.Second)
}
