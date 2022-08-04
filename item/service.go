package item

import "github.com/motikingo/ecommerceRESTAPI-Go/entity"

type ItemService interface{
	GetItems()([]entity.Item,[]error)
	GetItem(id uint)(*entity.Item,[]error)
	UpdateItem(id uint,item entity.Item)(*entity.Item,[]error)
	CreateItem(item entity.Item)(*entity.Item,[]error)
	DeleteItem(id uint)(*entity.Item,[]error)
}