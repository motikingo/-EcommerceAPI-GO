package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	"github.com/motikingo/ecommerceRESTAPI-Go/customer"
)

type UserRepo struct{
	db * gorm.DB
}

func NewUserRepo(db * gorm.DB) user.UserRepository{
	return &UserRepo{db:db}
}

func(usrRepo *UserRepo) GetUsers()([]entity.Customer,[]error){

	var users []entity.Customer 
	errs := usrRepo.db.Find(&users).GetErrors()
	if len(errs)>0 {
		return nil,errs
	}
	return users,nil
}

func(usrRepo *UserRepo)GetUser(id uint)(*entity.Customer,[]error){

	var user entity.Customer
	errs := usrRepo.db.First(&user,id).GetErrors()
	if len(errs)>0 {
		return nil,errs
	}
	return &user,nil
}

func (usrRepo *UserRepo) GetUserByUserName(name string) *entity.Customer {
	var user entity.Customer
	errs := usrRepo.db.First(&user,name).GetErrors()	
	if len(errs)>0{
		return nil
	}
	return &user 
}


func (usrRepo *UserRepo)GetUserByEmail(email string)bool{
	var user entity.Customer
	errs := usrRepo.db.First(&user,email).GetErrors()	
	return errs ==nil
}

func(usrRepo *UserRepo) CreateUser(user entity.Customer)(*entity.Customer,[]error){
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

func(usrRepo *UserRepo) UpdateUser(user entity.Customer)(*entity.Customer,[]error){
	usr,ers := usrRepo.GetUser(user.ID) 

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

func(usrRepo *UserRepo) DeleteUser(id uint)(*entity.Customer,[]error){
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


