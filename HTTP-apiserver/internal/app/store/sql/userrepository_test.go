package sql_test

import (
	"database/sql"
	"testing"

	"github.com/n1ghtwell/API-Server-Go/internal/app/model"
	"github.com/n1ghtwell/API-Server-Go/internal/app/store/sql"
	"github.com/stretchr/testify/assert"
)

func TestUserRepsoitory_Create(t *testing.T) {
	db, teardown := sql.TestDB(t, databaseURL)
	defer teardown("users")

	s := sql.New(db)
	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRespository_FIndByEmail(t *testing.T) {
	db, teardown := sql.TestDB(t, databaseURL)
	defer teardown("users")

	s := sql.New(db)
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
