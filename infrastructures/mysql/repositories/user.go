package repositories

import (
	"in-gravity/domains/entities"
	"in-gravity/domains/repositories"
)

type MySQLRepository struct {
}

func NewUserRepository() repositories.UserRepositoryInterface {
	return MySQLRepository{}
}

func (r MySQLRepository) Save(user entities.User) error {
	return nil
}
