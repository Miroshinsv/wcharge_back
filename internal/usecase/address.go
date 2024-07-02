package usecase

import (
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"log"
)

func (uc *UseCase) CreateAddress(address entity.Address) (*entity.Address, error) {
	adrs, err := uc.postgres.CreateAddress(address)
	if err != nil {
		log.Printf("Error - UseCase - CreateAddress - CreateAddressRepo: %s", err)
		return nil, err
	}

	return adrs, nil
}

func (uc *UseCase) UpdateAddress(id int, address entity.Address) (*entity.Address, error) {
	adrs, err := uc.postgres.UpdateAddress(address, id)
	if err != nil {
		log.Printf("Error - UseCase - CreateAddress - UpdateAddressRepo: %s", err)
		return nil, err
	}

	return adrs, nil
}

func (uc *UseCase) DeleteAddress(id int) error {
	err := uc.postgres.DeleteAddress(id)
	if err != nil {
		log.Printf("Error - UseCase - CreateAddress - DeleteAddressRepo: %s", err)
		return err
	}

	return nil
}

func (uc *UseCase) GetAddress(id int) (*entity.Address, error) {
	addr, err := uc.postgres.GetAddress(id)
	if err != nil {
		log.Printf("Error - UseCase - CreateAddress - GetAddress: %s", err)
		return nil, err
	}

	return addr, nil
}

func (uc *UseCase) GetAddresses() (*[]entity.Address, error) {
	addresses, err := uc.postgres.GetAddresses()
	if err != nil {
		log.Printf("Error - UseCase - CreateAddress - GetAddresses: %s", err)
		return nil, err
	}

	return addresses, nil
}
