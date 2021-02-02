package model_test

import (
	"testing"

	"github.com/codeedu/imersao/codepix-go/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func testNewUser(t *testing.T) {
	name := "Taco"
	email := "tacobarry@gmail.com"
	user, err := model.NewUser(name, email)

	require.Nil(t, err)
	require.NotNil(t, uuid.FromStringOrNil(user.ID))
	require.NotNil(t, uuid.FromStringOrNil(user.Name))
	require.NotNil(t, uuid.FromStringOrNil(user.Email))
}