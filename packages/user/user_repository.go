package user

import "restful_api/entities"

// UserRepository interface
type UserRepository interface {
	FetchAll() ([]*entities.User, error)
	FindByID(id uint) (*entities.User, error)
	Store(u *entities.User) (bool, error)
	Update(u *entities.User) (bool, error)
	Delete(id uint) (bool, error)
}
