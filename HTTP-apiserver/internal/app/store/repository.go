package store

import "github.com/n1ghtwell/API-Server-Go/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
