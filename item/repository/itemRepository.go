package itemRepository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	"github.com/motikingo/ecommerceRESTAPI-Go/item"
)

type ItemRepo struct {
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) item.ItemRepository{
	return &ItemRepo{db: db}
}

func(itemRepo *ItemRepo) GetItems()([]entity.Item,[]error){
	var items []entity.Item
	ers := itemRepo.db.Find(&items).GetErrors()
	if len(ers)>0{
		return nil,ers
	}
	return items,nil
}
func(itemRepo *ItemRepo)GetItem(id uint)(*entity.Item,[]error){

	var item entity.Item
	ers := itemRepo.db.First(&item,id).GetErrors()
	if len(ers)>0{
		return nil,ers
	}
	return &item,nil
}

func(itemRepo *ItemRepo)IsItemNameExist(name string) *entity.Item{

	var item entity.Item
	ers := itemRepo.db.First(&item,name).GetErrors()
	if len(ers)>0{
		return nil
	}
	return &item
}
func(itemRepo *ItemRepo)UpdateItem(item entity.Item)(*entity.Item,[]error){

	itm,e := itemRepo.GetItem(item.ID)

	if len(e)>0{
		return nil,e
	}
	itm.Name = item.Name
	itm.Description = item.Description
	itm.Brand = item.Brand
	itm.Price = item.Price
	itm.Image = item.Image
	itm.ProductionDate = item.ProductionDate
	itm.ExpireDate = item.ExpireDate
	
	ers := itemRepo.db.Save(&item).GetErrors()
	if len(ers)>0{
		return nil,ers
	}
	return &item,nil
}
func(itemRepo *ItemRepo)CreateItem(item entity.Item)(*entity.Item,[]error){
	itm ,ers:= itemRepo.GetItem(item.ID)
	if itm !=nil && len(ers)==0{
		log.Fatal("Item with this Id already exist")
		return nil,nil
	}else if itm==nil && len(ers)>0{
		ers := itemRepo.db.Create(&itm).GetErrors()
		if len(ers)>0{
			return nil,ers
		}
		return &item,nil
	}
	return nil,nil
	
}
func(itemRepo *ItemRepo)DeleteItem(id uint)(*entity.Item,[]error){
	
	item,ers:= itemRepo.GetItem(id)
	if item==nil && len(ers)>0{
		log.Fatal("this user doesn't exist")
		return nil,ers
	}else if item!=nil && len(ers)==0{
		ers := itemRepo.db.Delete(id).GetErrors()
		if len(ers)>0{
			return nil,ers
		}
		return item,nil
	}

	return nil,nil

}