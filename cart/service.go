package cart

import "github.com/motikingo/ecommerceRESTAPI-Go/entity"

type CartService interface{
	GetCarts()([]entity.Cart,[]error)
	GetCart(id uint)(*entity.Cart,[]error)
	GetCartByUserID(user_Id uint)*entity.Cart
	UpdateCart(id uint,car entity.Cart)(*entity.Cart,[]error)
	CreateCart(car entity.Cart)(*entity.Cart,[]error)
	DeleteCart(id uint)(*entity.Cart,[]error)

}
