package catagoryRepository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	"github.com/motikingo/ecommerceRESTAPI-Go/catagory"
)


type CatRepo struct{
	db *gorm.DB
}
func NewCatagoryRepo(db *gorm.DB) catagory.CatagoryRepository{
	return &CatRepo{db:db}
}

func(usrRepo *CatRepo) GetCatagories() ([]entity.Catagory,[]error){

	var cats []entity.Catagory
	ers:= usrRepo.db.Find(&cats).GetErrors()
	if len(ers)>0{
		return nil,ers
	}
	return cats,nil
}
func(usrRepo *CatRepo)GetCatagory(id uint) (*entity.Catagory,[]error){

	var cat entity.Catagory

	ers:= usrRepo.db.First(&cat,id).GetErrors()
	if len(ers)>0{
		return nil,ers
	}
	return &cat,nil
}

func(usrRepo *CatRepo)IsCatagoryNameExist(name string) bool{
	var cat entity.Catagory

	ers:= usrRepo.db.First(&cat,name).GetErrors()
	if len(ers)>0{
		return false
	}
	return true

}
func(usrRepo *CatRepo)UpdateCatagory(ct entity.Catagory) (*entity.Catagory,[]error){


	cat,ers := usrRepo.GetCatagory(ct.ID)
	if len(ers)>0{
		return nil,ers
	}
	cat.Name = ct.Name
	cat.Description = ct.Description
	cat.Image = ct.Image
	cat.Items_Id = ct.Items_Id 

	ers = usrRepo.db.Save(&cat).GetErrors()
	if len(ers)>0{
		return nil,ers
	}
	return cat,nil

}
func(usrRepo *CatRepo)CreateCatagory(ct entity.Catagory) (*entity.Catagory,[]error){

	
	cat,ers := usrRepo.GetCatagory(ct.ID)
	
	if len(ers)>0{
		return nil,ers
	}

	if cat !=nil && len(ers)==0 {
		log.Fatal("user already exist with this ID")
		return nil,nil
	}else if cat ==nil && len(ers)>0{
		cat.Name = ct.Name
		cat.Description = ct.Description
		cat.Items_Id = ct.Items_Id 

		ers = usrRepo.db.Save(&cat).GetErrors()
		if len(ers)>0{
			return nil,ers
		}
		return cat,nil
	}else{
		log.Fatal("what the heck is this")
	}
	return nil,nil

	

}
func(usrRepo *CatRepo)DeleteCatagory(id uint) (*entity.Catagory,[]error){

	cat,ers := usrRepo.GetCatagory(id)
	
	if len(ers)>0{
		return nil,ers
	}

	if cat ==nil && len(ers)>=0 {
		log.Fatal("user doesn't exist with this ID")
		return nil,nil
	}else if cat != nil && len(ers)>0{

		ers = usrRepo.db.Delete(&cat).GetErrors()
		if len(ers)>0{
			return nil,ers
		}
		return cat,nil
	}else{
		log.Fatal("what the heck is this")
	}
	return nil,nil
}

