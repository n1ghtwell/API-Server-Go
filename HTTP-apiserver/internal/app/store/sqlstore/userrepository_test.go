package sqlstore_test

import (
	"testing"

	"github.com/n1ghtwell/API-Server-Go/internal/app/model"
	"github.com/n1ghtwell/API-Server-Go/internal/app/store"
	"github.com/n1ghtwell/API-Server-Go/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepsoitory_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRespository_FIndByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
