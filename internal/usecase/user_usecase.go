package usecase

import (
	"github.com/faizinkholiq/go-clean-arch_boilerplate/internal/domain"
	"github.com/faizinkholiq/go-clean-arch_boilerplate/internal/repository"
)

type UserUseCase struct {
	UserRepo *repository.UserRepository
}

func NewUserUseCase(repo *repository.UserRepository) *UserUseCase {
	return &UserUseCase{UserRepo: repo}
}

func (uc *UserUseCase) GetUserList() ([]domain.User, error) {
	return uc.UserRepo.FindAll()
}

func (uc *UserUseCase) GetUserByID(id int) (*domain.User, error) {
	return uc.UserRepo.FindByID(id)
}

func (uc *UserUseCase) CreateUser(user domain.User) error {
	return uc.UserRepo.Save(user)
}
