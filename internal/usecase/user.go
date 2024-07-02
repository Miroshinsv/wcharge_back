package usecase

import (
	"log"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

func (uc *UseCase) CreateUser(u entity.User) (*entity.User, error) {
	uu, err := uc.postgres.CreateUser(u)
	if err != nil {
		log.Printf("Error - UserUseCase - CreateUser - uc.repo.CreateUser: %s", err)
		return nil, err
	}

	return uu, nil
}

func (uc *UseCase) UpdateUser(u entity.User, id int) (*entity.User, error) {
	uu, err := uc.postgres.UpdateUser(u, id)
	if err != nil {
		log.Printf("Error - UserUseCase - UpdateUser - uc.repo.UpdateUser: %s", err)
		return nil, err
	}

	return uu, nil
}

func (uc *UseCase) DeleteUser(id int) error {
	err := uc.postgres.DeleteUser(id)
	if err != nil {
		log.Printf("Error - UserUseCase - DeleteUser - uc.repo.DeleteUser: %s", err)
		return err
	}

	return nil
}

func (uc *UseCase) GetUser(id int) (*entity.User, error) {
	user, err := uc.postgres.GetUser(id)
	if err != nil {
		log.Printf("Error - UserUseCase - GetUsers - uc.repo.GetUsers: %s", err)
		return nil, err
	}

	return user, nil
}

func (uc *UseCase) GetUsers() (*[]entity.User, error) {
	users, err := uc.postgres.GetUsers()
	if err != nil {
		log.Printf("Error - UserUseCase - GetUsers - uc.repo.GetUsers: %s", err)
		return nil, err
	}

	return users, nil
}

func (uc *UseCase) GetUserByName(userName string) (*entity.User, error) {
	u, err := uc.postgres.GetUserByName(userName)
	if err != nil {
		log.Printf("Error - UseCase - GetUserByName - uc.repo.GetUserByName: %s", err)
		return nil, err
	}

	return u, nil
}
