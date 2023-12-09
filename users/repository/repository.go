package repository

import "github.com/victoorraphael/simple-ms/users/entities"

type UserRepo interface {
	List() ([]entities.User, error)
	Get(id int64) (entities.User, error)
	Create(user entities.User) error
	Update(user entities.User) error
	Delete(id int64) error
}
