package usecase

import (
	"fmt"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

func (uc *UseCase) CreateUser(u entity.User) error {
	err := uc.postgres.CreateUserRepo(u)
	if err != nil {
		return fmt.Errorf("UserUseCase - CreateUser - uc.repo.CreateUser: %w", err)
	}

	return nil
}

func (uc *UseCase) UpdateUser(u entity.User) error {
	err := uc.postgres.UpdateUserRepo(u)
	if err != nil {
		return fmt.Errorf("UserUseCase - UpdateUser - uc.repo.UpdateUser: %w", err)
	}

	return nil
}

func (uc *UseCase) DeleteUser(u entity.User) error {
	err := uc.postgres.DeleteUserRepo(u)
	if err != nil {
		return fmt.Errorf("UserUseCase - DeleteUser - uc.repo.DeleteUser: %w", err)
	}

	return nil
}

func (uc *UseCase) GetUser(u entity.User) (entity.User, error) {
	user, err := uc.postgres.GetUserRepo(u)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserUseCase - GetUsers - uc.repo.GetUsers: %w", err)
	}

	return user, nil
}

func (uc *UseCase) GetUsers() ([]entity.User, error) {
	users, err := uc.postgres.GetUsersRepo()
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - GetUsers - uc.repo.GetUsers: %w", err)
	}

	return users, nil
}
