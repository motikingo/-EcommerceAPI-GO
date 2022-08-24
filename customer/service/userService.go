package service

import (
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	"github.com/motikingo/ecommerceRESTAPI-Go/customer"
)

type UserSrvc struct{
	repo user.UserRepository
}

func NewUserSrvc(repo user.UserRepository) user.UserService{
	return &UserSrvc{repo:repo}
}

func(usrRepo *UserSrvc) GetUsers()([]entity.Customer,[]error){
	users,errs:= usrRepo.repo.GetUsers()

	if  len(errs)>0 {
		return nil,errs
	}

	return users,nil

}


func(usrRepo *UserSrvc)  GetUser(id uint)(*entity.Customer,[]error){
	
	user,errs:= usrRepo.repo.GetUser(id)

	if  len(errs)>0 {
		return nil,errs
	}

	return user,nil
}

func(usrRepo *UserSrvc)GetUserByUserName(name string) *entity.Customer{
	user := usrRepo.repo.GetUserByUserName(name)
	return user
}
func(usrRepo *UserSrvc)GetUserByEmail(email string)bool{	
	return usrRepo.repo.GetUserByEmail(email)
}

func(usrRepo *UserSrvc) CreateUser(user entity.Customer)(*entity.Customer,[]error){
	
	usr,errs:= usrRepo.repo.CreateUser(user)

	if  len(errs)>0 {
		return nil,errs
	}

	return usr,nil
}

func(usrRepo *UserSrvc) UpdateUser(user entity.Customer)(*entity.Customer,[]error){
	usr,errs:= usrRepo.repo.UpdateUser(user)

	if  len(errs)>0 {
		return nil,errs
	}

	return usr,nil
}

func(usrRepo *UserSrvc) DeleteUser(id uint)(*entity.Customer,[]error){
	user,errs:= usrRepo.repo.DeleteUser(id)

	if  len(errs)>0 {
		return nil,errs
	}

	return user,nil
}