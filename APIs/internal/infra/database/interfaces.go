package database

import "github.com/marcofilho/Pos-GO-Expert/APIs/internal/entity"

type UserDBInterface interface {
	CreateUser(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
