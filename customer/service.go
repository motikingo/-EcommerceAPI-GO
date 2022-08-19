package user

import (
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
)

type UserService interface{
	GetUsers()([]entity.User,[]error)
	GetUser(id uint)(*entity.User,[]error)
	CreateUser(user entity.User)(*entity.User,[]error)
	UpdateUser(id uint,user entity.User)(*entity.User,[]error)
	DeleteUser(id uint)(*entity.User,[]error)

}