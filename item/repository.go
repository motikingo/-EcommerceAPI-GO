package item

import "github.com/motikingo/ecommerceRESTAPI-Go/entity"

type ItemRepository interface{
	GetItems()([]entity.Item,[]error)
	GetItem(id uint)(*entity.Item,[]error)
	IsItemNameExist(name string) *entity.Item
	UpdateItem(item entity.Item)(*entity.Item,[]error)
	CreateItem(item entity.Item)(*entity.Item,[]error)
	DeleteItem(id uint)(*entity.Item,[]error)
}
