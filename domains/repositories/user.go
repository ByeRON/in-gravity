package repositories

import "in-gravity/domains/entities"

type UserRepositoryInterface interface {
	Save(entities.User) error
}
