package repository

import (
	"database/sql"

	"github.com/faizinkholiq/go-clean-arch_boilerplate/internal/domain"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	return users, nil
}

func (repo *UserRepository) FindByID(id int) (*domain.User, error) {
	var user domain.User
	err := repo.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) Save(user domain.User) error {
	_, err := repo.DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
		user.Name, user.Email, user.Password)
	return err
}
