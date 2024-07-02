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
		CreateUser(entity.User) (*entity.User, error)
		UpdateUser(entity.User, int) (*entity.User, error)
		DeleteUser(int) error
		GetUser(int) (*entity.User, error)
		GetUsers() (*[]entity.User, error)
		GetUserByName(userName string) (*entity.User, error)
	}

	UserRepo interface {
		GetUserByName(string) (*entity.User, error)
		CreateUser(entity.User) (*entity.User, error)
		UpdateUser(entity.User, int) (*entity.User, error)
		DeleteUser(int) error
		GetUser(int) (*entity.User, error)
		GetUsers() (*[]entity.User, error)
	}
)

type (
	StationAPI interface {
		CreateStation(entity.Station) (*entity.Station, error)
		UpdateStation(int, entity.Station) (*entity.Station, error)
		DeleteStation(int) error
		GetStation(int) (*entity.Station, error)
		GetStations() (*[]entity.Station, error)
	}

	StationRepo interface {
		CreateStation(entity.Station) (*entity.Station, error)
		UpdateStation(entity.Station, int) (*entity.Station, error)
		DeleteStation(int) error
		GetStation(int) (*entity.Station, error)
		GetStations() (*[]entity.Station, error)
	}
)

type (
	PowerbankAPI interface {
		CreatePowerbank(entity.Powerbank) (*entity.Powerbank, error)
		UpdatePowerbank(int, entity.Powerbank) (*entity.Powerbank, error)
		DeletePowerbank(int) error
		GetPowerbank(int) (*entity.Powerbank, error)
		GetPowerbanks() (*[]entity.Powerbank, error)
		GetAllPowerbanksInStation(int) (*[]entity.Powerbank, error)
	}

	PowerbankRepo interface {
		CreatePowerbank(entity.Powerbank) (*entity.Powerbank, error)
		UpdatePowerbank(entity.Powerbank, int) (*entity.Powerbank, error)
		DeletePowerbank(int) error
		GetRamdomPowerbank() (*entity.Powerbank, error)
		GetPowerbank(int) (*entity.Powerbank, error)
		GetPowerbanks() (*[]entity.Powerbank, error)
		GetAllPowerbanksInStation(int) (*[]entity.Powerbank, error)
	}
)

type (
	AddressAPI interface {
		CreateAddress(entity.Address) (*entity.Address, error)
		UpdateAddress(int, entity.Address) (*entity.Address, error)
		DeleteAddress(int) error
		GetAddress(int) (*entity.Address, error)
		GetAddresses() (*[]entity.Address, error)
	}

	AddressRepo interface {
		CreateAddress(entity.Address) (*entity.Address, error)
		UpdateAddress(entity.Address, int) (*entity.Address, error)
		DeleteAddress(int) error
		GetAddress(int) (*entity.Address, error)
		GetAddresses() (*[]entity.Address, error)
	}
)

type (
	RoleRepo interface {
		GetRole(int) (*entity.Role, error)
		GetRoles() (*[]entity.Role, error)
	}
)

type (
	UserActionAPI interface {
		GetUserPowerbanks(int) (*[]entity.Powerbank, error)
		TakePowerbank(int, int) (*entity.Powerbank, error)
		PutPowerbank(int, int, int, int) error
		AddPowerbankToStation(int, int, int) error
	}

	UserActionRepo interface {
		GetUsersPowerbanks(int) (*[]entity.Powerbank, error)
		InsertStationPowerbank(int, int, int) error
		DeleteStationPowerbank(int, int) error
		InsertUserPowerbank(int, int, int) error
		DeleteUserPowerbank(int, int) error
		TakePowerbank(int, int) error
		PutPowerbank(int, int, int, int) error
		AddPowerbankToStation(int, int, int) error
	}
)
