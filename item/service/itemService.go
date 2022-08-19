package itemService

import (
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	"github.com/motikingo/ecommerceRESTAPI-Go/item"
)

type ItemServ struct {
	repo item.ItemRepository
}

func NewItemServ(repo item.ItemRepository) item.ItemService{
	return &ItemServ{repo: repo}
}

func(itemsrv *ItemServ) GetItems()([]entity.Item,[]error){
	items,ers:= itemsrv.repo.GetItems()
	if len(ers)>0{
		return nil,ers
	}
	return items,nil

}
func(itemsrv *ItemServ)GetItem(id uint)(*entity.Item,[]error){

	item,ers:= itemsrv.repo.GetItem(id)
	if len(ers)>0{
		return nil,ers
	}
	return item,nil
}
func(itemsrv *ItemServ)UpdateItem(item entity.Item)(*entity.Item,[]error){

	itm,ers:= itemsrv.repo.UpdateItem(id,item)
	if len(ers)>0{
		return nil,ers
	}
	return itm,nil
}
func(itemsrv *ItemServ)CreateItem(item entity.Item)(*entity.Item,[]error){

	items,ers:= itemsrv.repo.CreateItem(item)
	if len(ers)>0{
		return nil,ers
	}
	return items,nil
}
func(itemsrv *ItemServ)DeleteItem(id uint)(*entity.Item,[]error){
	items,ers:= itemsrv.repo.DeleteItem(id)
	if len(ers)>0{
		return nil,ers
	}
	return items,nil
}