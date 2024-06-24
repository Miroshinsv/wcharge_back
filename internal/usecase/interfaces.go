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

	// TODO
	UserRepo interface {
		GetUserByNameRepo(userName string) (*entity.User, error)
		CreateUserRepo(u entity.User) error
		UpdateUserRepo(u entity.User, id int) error
		DeleteUserRepo(id int) error
		GetUserRepo(id int) (*entity.User, error)
		GetUsersRepo() (*[]entity.User, error)
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
		CreateStation(s entity.Station) (*entity.Station, error)
		UpdateStation(s entity.Station, id int) (*entity.Station, error)
		DeleteStationRepo(id int) error
		GetStation(id int) (*entity.Station, error)
		GetStations() (*[]entity.Station, error)
		//GetAllPowerbanksInStation(stationId int) (*[]entity.Powerbank, error) // TODO move to powerbanks
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
		CreatePowerbank(p entity.Powerbank) (*entity.Powerbank, error)
		UpdatePowerbank(p entity.Powerbank, id int) (*entity.Powerbank, error)
		DeletePowerbank(id int) error
		GetRamdomPowerbank() (*entity.Powerbank, error)
		GetPowerbank(id int) (*entity.Powerbank, error)
		GetPowerbanks() (*[]entity.Powerbank, error)
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
		CreateAddress(addr entity.Address) (*entity.Address, error)
		UpdateAddress(addr entity.Address, id int) (*entity.Address, error)
		DeleteAddress(id int) error
		GetAddress(id int) (*entity.Address, error)
		GetAddresses() (*[]entity.Address, error)
	}
)

type (
	RoleAPI interface {
		GetRole(int) (entity.Role, error)
		GetRoles() ([]entity.Role, error)
	}

	RoleRepo interface {
		GetRole(id int) (*entity.Role, error)
		GetRoles() (*[]entity.Role, error)
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
		GetUserPowerbanksRepo(userId int) (*[]entity.Powerbank, error)
		InsertStationPowerbank(powerbankId int, stationId int, position int) error
		TakePowerbank(userId int, powerbankId int) error
		PutPowerbankRepo(userId int, powerbankId int, stationId int, position int) error
		AddPowerbankToStationRepo(powerbankId int, stationId int, position int) error
	}
)
