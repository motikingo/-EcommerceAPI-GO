package cart

import "github.com/motikingo/ecommerceRESTAPI-Go/entity"

type CartRepository interface{
	GetCarts()([]entity.Cart,[]error)
	GetCart(id uint)(*entity.Cart,[]error)
	GetCartByUserID(user_Id uint)*entity.Cart
	UpdateCart(car entity.Cart)(*entity.Cart,[]error)
	CreateCart(car entity.Cart)(*entity.Cart,[]error)
	DeleteCart(id uint)(*entity.Cart,[]error)

}
