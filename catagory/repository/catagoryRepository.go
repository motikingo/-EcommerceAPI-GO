package catagoryRepository

import (
	//"fmt"
	"log"
	"reflect"

	"github.com/jinzhu/gorm"
	"github.com/motikingo/ecommerceRESTAPI-Go/catagory"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
)


type CatRepo struct{
	db *gorm.DB
}
func NewCatagoryRepo(db *gorm.DB) catagory.CatagoryRepository{
	return &CatRepo{db:db}
}

func(usrRepo *CatRepo) GetCatagories() ([]entity.Catagory,[]error){

	var cats []entity.Catagory
	ers:= usrRepo.db.Preload("Items").Find(&cats).GetErrors()
	if len(ers)>0{
		return nil,ers
	}
	return cats,nil
}
func(usrRepo *CatRepo)GetCatagory(id uint) (*entity.Catagory,[]error){

	var cat entity.Catagory

	ers:= usrRepo.db.Preload("Items").First(&cat,id).GetErrors()
	if len(ers)>0{
		return nil,ers
	}
	return &cat,nil
}

func(usrRepo *CatRepo)IsCatagoryNameExist(name string) bool{
	var cat entity.Catagory
	if ers:= usrRepo.db.Where("name=?",name).First(&cat).GetErrors(); len(ers)>0{
		return false
	}
	return true

}
func(usrRepo *CatRepo)CreateCatagory(ct entity.Catagory) (*entity.Catagory,[]error){

	cat,ers := usrRepo.GetCatagory(ct.ID)
	
	if cat !=nil && len(ers)==0 {
		log.Fatal("user already exist with this ID")
		return cat,nil
	}
	if er := usrRepo.db.Create(&ct).GetErrors(); len(er)>0{
		return nil,er
	}
	return cat,nil
 
}

func(usrRepo *CatRepo)UpdateCatagory(ct entity.Catagory) (*entity.Catagory,[]error){

	cat,ers := usrRepo.GetCatagory(ct.ID)
	if len(ers)>0{
		return nil,ers
	}
	cat.Name = func () string {
		if ct.Name != cat.Name{
			return ct.Name
		}
		return cat.Name
	}()
	cat.Description = func () string {
		if ct.Description != cat.Description{
			return ct.Description
		}
		return cat.Description
	}()
	cat.Image = func () string  {
		if ct.Image != cat.Image{
			return ct.Image
		}
		return cat.Image
	}()
	cat.Items = func () []entity.Item {
		if !reflect.DeepEqual(ct.Items,cat.Items){
			usrRepo.db.Model(&cat).Association("Items").Clear()
			return ct.Items 
		} 
		return cat.Items 
	}()

	ers = usrRepo.db.Save(&cat).GetErrors()
	if len(ers)>0{
		return nil,ers
	}
	return cat,nil

}
func(usrRepo *CatRepo)DeleteCatagory(id uint) (*entity.Catagory,[]error){

	cat,ers := usrRepo.GetCatagory(id)
	
	if len(ers)>0{
		return nil,ers
	}

	if cat ==nil && len(ers)>=0 {
		log.Fatal("user doesn't exist with this ID")
		return nil,nil
	}else if cat != nil && len(ers)==0{

		if ers = usrRepo.db.Delete(&cat).GetErrors(); len(ers)>0{
			return nil,ers
		}
		usrRepo.db.Model(cat).Association("Items").Clear()
		return cat,nil
	}else{
		log.Fatal("what the heck is this")
	}
	return nil,nil
}

