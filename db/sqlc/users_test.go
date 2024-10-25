package db

import (
	"context"
	"expenses/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomUser() (User, error) {
	arg := CreateUserParams{
		Username:       util.RandomUsername(),
		Email:          util.RandomEmail(),
		FullName:       util.RandomUsername(),
		HashedPassword: "secret",
	}

	user, err := queriesTest.CreateUser(context.Background(), arg)

	if arg.Username != user.Username {
		return user, err
	}
	if arg.Email != user.Email {
		return user, err
	}
	if arg.FullName != user.FullName {
		return user, err
	}

	return user, nil
}

func TestCreateRandomUser(t *testing.T) {
	user, err := CreateRandomUser()
	require.NoError(t, err)
	require.NotEmpty(t, user)
}
