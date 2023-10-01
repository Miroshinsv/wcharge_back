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
		PowerbankRepo
	}
)

type (
	UserAPI interface {
		Login(entity.UserLogin) error
		CreateUser(entity.User) error
		UpdateUser(int, entity.User) error
		DeleteUser(int) error
		GetUser(int) (entity.User, error)
		GetUsers() ([]entity.User, error)
		GetUserByName(userName string) (entity.User, error)
	}

	UserRepo interface {
		CreateUserRepo(entity.User) error
		UpdateUserRepo(int, entity.User) error
		DeleteUserRepo(int) error
		GetUserRepo(int) (entity.User, error)
		GetUsersRepo() ([]entity.User, error)
		GetUserByNameRepo(userName string) (entity.User, error)
	}
)

type (
	StationAPI interface {
		CreateStation(station entity.Station) error
		UpdateStation(int, entity.Station) error
		DeleteStation(int) error
		GetStation(int) (entity.Station, error)
		GetStations() ([]entity.Station, error)
	}

	StationRepo interface {
		CreateStationRepo(entity.Station) error
		UpdateStationRepo(int, entity.Station) error
		DeleteStationRepo(int) error
		GetStationRepo(int) (entity.Station, error)
		GetStationsRepo() ([]entity.Station, error)
	}
)

type (
	PowerbankAPI interface {
		CreatePowerbank(station entity.Powerbank) error
		UpdatePowerbank(int, entity.Powerbank) error
		DeletePowerbank(int) error
		GetPowerbank(int) (entity.Powerbank, error)
		GetPowerbanks() ([]entity.Powerbank, error)
	}

	PowerbankRepo interface {
		CreatePowerbankRepo(entity.Powerbank) error
		UpdatePowerbankRepo(int, entity.Powerbank) error
		DeletePowerbankRepo(int) error
		GetPowerbankRepo(int) (entity.Powerbank, error)
		GetPowerbanksRepo() ([]entity.Powerbank, error)
	}
)
