package catagory

import "github.com/motikingo/ecommerceRESTAPI-Go/entity"

type CatagoryRepository interface {
	GetCatagories() ([]entity.Catagory,[]error)
	GetCatagory(id uint) (*entity.Catagory,[]error)
	IsCatagoryNameExist(name string) bool
	UpdateCatagory(cat entity.Catagory) (*entity.Catagory,[]error)
	CreateCatagory(usr entity.Catagory) (*entity.Catagory,[]error)
	DeleteCatagory(id uint) (*entity.Catagory,[]error)

}
