package service

import (
	"github.com/motikingo/ecommerceRESTAPI-Go/catagory"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"

)

type UserServc struct {
	repo catagory.CatagoryRepository
}

func NewUsrRepo(repo catagory.CatagoryRepository) catagory.CatagoryService{ 
	return &UserServc{repo: repo}
}

func (catRepo *UserServc) GetCatagories() ([]entity.Catagory, []error) {

	cat,ers:= catRepo.repo.GetCatagories()
	if len(ers)>0{
		return nil,ers
	}
	return cat,nil
}
func (catRepo *UserServc) GetCatagory(id uint) (*entity.Catagory, []error) {

	cat,ers:= catRepo.repo.GetCatagory(id)
	if len(ers)>0{
		return nil,ers
	}
	return cat,nil

}
func (catRepo *UserServc) UpdateCatagory(id uint, ct entity.Catagory) (*entity.Catagory, []error) {

	cat,ers:= catRepo.repo.UpdateCatagory(id,ct)
	if len(ers)>0{
		return nil,ers
	}
	return cat,nil

}
func (catRepo *UserServc) CreateCatagory(ct entity.Catagory) (*entity.Catagory, []error) {

	cat,ers:= catRepo.repo.CreateCatagory(ct)
	if len(ers)>0{
		return nil,ers
	}
	return cat,nil
}
func (catRepo *UserServc) DeleteCatagory(id uint) (*entity.Catagory, []error) {

	cat,ers:= catRepo.repo.DeleteCatagory(id)
	if len(ers)>0{
		return nil,ers
	}
	return cat,nil
}