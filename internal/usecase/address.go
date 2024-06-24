package usecase

import (
	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

// TODO
func (uc *UseCase) CreateAddress(address entity.Address) error {
	//err := uc.postgres.CreateAddress(address)
	//if err != nil {
	//	return fmt.Errorf("UseCase - CreateAddress - CreateAddressRepo - " + err.Error())
	//}
	return nil
}

func (uc *UseCase) UpdateAddress(id int, address entity.Address) error {
	//err := uc.postgres.UpdateAddressRepo(id, address)
	//if err != nil {
	//	return fmt.Errorf("UseCase - CreateAddress - UpdateAddressRepo - " + err.Error())
	//}
	return nil
}

func (uc *UseCase) DeleteAddress(id int) error {
	//err := uc.postgres.DeleteAddressRepo(id)
	//if err != nil {
	//	return fmt.Errorf("UseCase - CreateAddress - DeleteAddressRepo - " + err.Error())
	//}
	return nil
}

func (uc *UseCase) GetAddress(id int) (entity.Address, error) {
	//addr, err := uc.postgres.GetAddressRepo(id)
	//if err != nil {
	//	return entity.Address{}, fmt.Errorf("UseCase - CreateAddress - GetAddress - " + err.Error())
	//}

	return entity.Address{}, nil
}

func (uc *UseCase) GetAddresses() ([]entity.Address, error) {
	//addresses, err := uc.postgres.GetAddressesRepo()
	//if err != nil {
	//	return nil, fmt.Errorf("UseCase - CreateAddress - GetAddresses - " + err.Error())
	//}

	return nil, nil
}
