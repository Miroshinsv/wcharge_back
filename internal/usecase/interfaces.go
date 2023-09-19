// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Translation -.
	User interface {
		Translate(context.Context, entity.User) (entity.User, error)
		History(context.Context) ([]entity.User, error)
	}

	// TranslationRepo -.
	UserRepo interface {
		Store(context.Context, entity.User) error
		GetHistory(context.Context) ([]entity.User, error)
	}

	// TranslationWebAPI -.
	UserWebAPI interface {
		Translate(entity.User) (entity.User, error)
	}
)
