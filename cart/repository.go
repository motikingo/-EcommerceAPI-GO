package cart

import "github.com/motikingo/ecommerceRESTAPI-Go/entity"

type CartRepository interface{
	GetCarts()([]entity.Cart,[]error)
	GetCart(id uint)(*entity.Cart,[]error)
	UpdateCart(id uint,car entity.Cart)(*entity.Cart,[]error)
	CreateCart(car entity.Cart)(*entity.Cart,[]error)
	DeleteCart(id uint)(*entity.Cart,[]error)

}
