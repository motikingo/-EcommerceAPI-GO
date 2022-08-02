package service

import (
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	"github.com/motikingo/ecommerceRESTAPI-Go/user"
)

type UserSrvc struct{
	repo user.UserRepository
}

func NewUserSrvc(repo user.UserRepository) user.UserService{
	return &UserSrvc{repo:repo}
}

func(usrRepo *UserSrvc) GetUsers()([]entity.User,[]error){
	users,errs:= usrRepo.repo.GetUsers()

	if  len(errs)>0 {
		return nil,errs
	}

	return users,nil

}

func(usrRepo *UserSrvc)  GetUser(id uint)(*entity.User,[]error){
	
	user,errs:= usrRepo.repo.GetUser(id)

	if  len(errs)>0 {
		return nil,errs
	}

	return user,nil
}

func(usrRepo *UserSrvc) CreateUser(user entity.User)(*entity.User,[]error){
	
	usr,errs:= usrRepo.repo.CreateUser(user)

	if  len(errs)>0 {
		return nil,errs
	}

	return usr,nil
}

func(usrRepo *UserSrvc) UpdateUser(id uint,user entity.User)(*entity.User,[]error){
	usr,errs:= usrRepo.repo.UpdateUser(id,user)

	if  len(errs)>0 {
		return nil,errs
	}

	return usr,nil
}

func(usrRepo *UserSrvc) DeleteUser(id uint)(*entity.User,[]error){
	user,errs:= usrRepo.repo.DeleteUser(id)

	if  len(errs)>0 {
		return nil,errs
	}

	return user,nil
}