package catagoryService

import (
	"github.com/motikingo/ecommerceRESTAPI-Go/catagory"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
)

type catagoryServc struct {
	repo catagory.CatagoryRepository
}

func NewCatagoryServ(repo catagory.CatagoryRepository) catagory.CatagoryService{ 
	return &catagoryServc{repo: repo}
}

func (catServc *catagoryServc) GetCatagories() ([]entity.Catagory, []error) {

	cat,ers:= catServc.repo.GetCatagories()
	if len(ers)>0{
		return nil,ers
	}
	return cat,nil
}
func (catServc *catagoryServc) GetCatagory(id uint) (*entity.Catagory, []error) {

	cat,ers:= catServc.repo.GetCatagory(id)
	if len(ers)>0{
		return nil,ers
	}
	return cat,nil

}
func (catServc *catagoryServc) UpdateCatagory(id uint, ct entity.Catagory) (*entity.Catagory, []error) {

	cat,ers:= catServc.repo.UpdateCatagory(id,ct)
	if len(ers)>0{
		return nil,ers
	}
	return cat,nil

}
func (catServc *catagoryServc) CreateCatagory(ct entity.Catagory) (*entity.Catagory, []error) {

	cat,ers:= catServc.repo.CreateCatagory(ct)
	if len(ers)>0{
		return nil,ers
	}
	return cat,nil
}
func (catServc *catagoryServc) DeleteCatagory(id uint) (*entity.Catagory, []error) {

	cat,ers:= catServc.repo.DeleteCatagory(id)
	if len(ers)>0{
		return nil,ers
	}
	return cat,nil
}