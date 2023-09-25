// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	PostgresRepo interface {
		UserRepo
		StationRepo
	}
)

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
	StationAPI interface {
		CreateStation(station entity.Station) error
		UpdateStation(entity.Station) error
		DeleteStation(entity.Station) error
		GetStation(entity.Station) (entity.Station, error)
		GetStations() ([]entity.Station, error)
	}

	StationRepo interface {
		CreateStationRepo(entity.Station) error
		UpdateStationRepo(entity.Station) error
		DeleteStationRepo(entity.Station) error
		GetStationRepo(entity.Station) (entity.Station, error)
		GetStationsRepo() ([]entity.Station, error)
	}
)
