package repositories

import (
	"database/sql"

	"github.com/faizinkholiq/go-clean-arch_boilerplate/internal/entities"
)

type UserRepo struct {
	DB *sql.DB
}

func (repo *UserRepo) CreateUser(user *entities.User) error {
	_, err := repo.DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
		user.Name, user.Email, user.Password)
	return err
}

func (repo *UserRepo) GetUserByID(id int) (*entities.User, error) {
	var user entities.User
	err := repo.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
