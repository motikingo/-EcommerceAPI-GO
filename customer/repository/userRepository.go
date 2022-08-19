package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	"github.com/motikingo/ecommerceRESTAPI-Go/user"
)

type UserRepo struct{
	db * gorm.DB
}

func NewUserRepo(db * gorm.DB) user.UserRepository{
	return &UserRepo{db:db}
}

func(usrRepo *UserRepo) GetUsers()([]entity.User,[]error){

	var users []entity.User
	errs := usrRepo.db.Find(&users).GetErrors()
	if len(errs)>0 {
		return nil,errs
	}
	return users,nil
}

func(usrRepo *UserRepo)  GetUser(id uint)(*entity.User,[]error){

	var user entity.User
	errs := usrRepo.db.First(&user,id).GetErrors()
	if len(errs)>0 {
		return nil,errs
	}
	return &user,nil
}

func(usrRepo *UserRepo) CreateUser(user entity.User)(*entity.User,[]error){
	usr,ers := usrRepo.GetUser(user.ID) 

	if usr!=nil && len(ers)==0 {

		log.Fatal("this user is already exist")
		return nil,ers
		
	}else if usr==nil && len(ers)>0{

		errs := usrRepo.db.Create(&user).GetErrors()
		if len(errs)>0 {
			return nil,errs
		}
		return &user,nil

		  
	}else{
		log.Fatal(ers)
	}
	return nil,nil
	
}

func(usrRepo *UserRepo) UpdateUser(id uint,user entity.User)(*entity.User,[]error){
	usr,ers := usrRepo.GetUser(id) 

	if usr!=nil && ers == nil {
		usr.Name = user.Name
		usr.LastName = user.LastName
		usr.UserName = user.UserName
		usr.Email = user.Email
		usr.Password = user.Password
		errs := usrRepo.db.Save(&usr).GetErrors()
		if len(errs)>0 {
			return nil,errs
		}
		return usr,nil

	}else if usr==nil && len(ers)>0{
		log.Fatal("this user deosn't is already exist creating user")
		urs,erU := usrRepo.CreateUser(*usr)
		if len(erU)>0 {
			return nil,erU
		}
		return urs,nil 
	}else{
		log.Fatal(ers)
	}
	return nil,nil
}

func(usrRepo *UserRepo) DeleteUser(id uint)(*entity.User,[]error){
	usr,ers := usrRepo.GetUser(id) 

	if usr!=nil && len(ers)==0 {

		errs := usrRepo.db.Delete(&usr).GetErrors()
		if len(errs)>0 {
			return nil,errs
		}
		return usr,nil
	}else if usr==nil && len(ers)>0{
		log.Fatal("this user is doesn't exist")
		return nil,ers  
	}else{
		log.Fatal(ers)
	}
	return nil,nil

}


