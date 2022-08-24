package user

import (
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
)

type UserService interface{
	GetUsers()([]entity.Customer,[]error)
	GetUser(id uint)(*entity.Customer,[]error)
	GetUserByUserName(name string) *entity.Customer
	GetUserByEmail(email string)bool
	CreateUser(user entity.Customer)(*entity.Customer,[]error)
	UpdateUser(user entity.Customer)(*entity.Customer,[]error)
	DeleteUser(id uint)(*entity.Customer,[]error)

}