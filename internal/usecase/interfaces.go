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
		AddressRepo
		RoleRepo
		UserActionRepo
	}
)

type (
	UserAPI interface {
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
		GetAllPowerbanksInStation(int) ([]entity.Powerbank, error)
	}

	StationRepo interface {
		CreateStationRepo(entity.Station) error
		UpdateStationRepo(int, entity.Station) error
		DeleteStationRepo(int) error
		GetStationRepo(int) (entity.Station, error)
		GetStationsRepo() ([]entity.Station, error)
		GetAllPowerbanksInStationRepo(int) ([]entity.Powerbank, error)
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

type (
	AddressAPI interface {
		CreateAddress(station entity.Address) error
		UpdateAddress(int, entity.Address) error
		DeleteAddress(int) error
		GetAddress(int) (entity.Address, error)
		GetAddresses() ([]entity.Address, error)
	}

	AddressRepo interface {
		CreateAddressRepo(entity.Address) error
		UpdateAddressRepo(int, entity.Address) error
		DeleteAddressRepo(int) error
		GetAddressRepo(int) (entity.Address, error)
		GetAddressesRepo() ([]entity.Address, error)
	}
)

type (
	RoleAPI interface {
		GetRole(int) (entity.Role, error)
		GetRoles() ([]entity.Role, error)
	}

	RoleRepo interface {
		GetRoleRepo(int) (entity.Role, error)
		GetRolesRepo() ([]entity.Role, error)
	}
)

type (
	UserActionAPI interface {
		GetUserPowerbanks(userId int) ([]entity.Powerbank, error) // all user's powerbanks
		TakePowerbank(userId int, powerbankId int, stationId int) error
		PutPowerbank(userId int, powerbankId int, stationId int) error
		AddPowerbankToStation(powerbankId int, stationId int) error
	}

	UserActionRepo interface {
		GetUserPowerbanksRepo(userId int) ([]entity.Powerbank, error) // all user's powerbanks
		TakePowerbankRepo(userId int, powerbankId int, stationId int) error
		PutPowerbankRepo(userId int, powerbankId int, stationId int) error
		AddPowerbankToStationRepo(powerbankId int, stationId int) error
	}
)
