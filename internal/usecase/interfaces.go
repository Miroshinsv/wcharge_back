// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	UserAPI interface {
		CreateUser(entity.User) error
		UpdateUser(entity.User) error
		DeleteUser(entity.User) error
		GetUser(entity.User) (entity.User, error)
		GetUsers() ([]entity.User, error)
	}

	UserRepo interface {
		CreateUserRepo(entity.User) error
		UpdateUserRepo(entity.User) error
		DeleteUserRepo(entity.User) error
		GetUserRepo(entity.User) (entity.User, error)
		GetUsersRepo() ([]entity.User, error)
	}
)

type (
	Powerbank interface {
	}

	PowerbankPostgres interface {
	}
)
