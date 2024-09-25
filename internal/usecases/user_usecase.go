package usecases

import "github.com/faizinkholiq/gofiber_boilerplate/internal/entities"

type UserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByID(id int) (*entities.User, error)
}

type UserUseCase struct {
	Repo UserRepository
}

func (uc *UserUseCase) RegisterUser(user *entities.User) error {
	return uc.Repo.CreateUser(user)
}
