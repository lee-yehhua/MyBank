package api

import (
	db "github.com/lee-yehhua/mybank/db/sqlc"
	"github.com/lee-yehhua/mybank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func randomUser(t *testing.T) (user db.User, password string) {
	password = util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user = db.User{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	return
}
