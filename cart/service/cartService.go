package cartService

import (
	"github.com/motikingo/ecommerceRESTAPI-Go/cart"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
)

type CartServ struct{
	repo cart.CartRepository
}

func NewCartServ(repo cart.CartRepository) cart.CartService{
	return &CartServ{repo:repo}
}

func(carServ *CartServ) GetCarts()([]entity.Cart,[]error){

	//var carts []entity.Cart
	carts,err := carServ.repo.GetCarts()
	if len(err)>0 {
		return nil,err
	}
	return carts,err

}
func(carServ *CartServ)GetCart(id uint)(*entity.Cart,[]error){

	//var cart entity.Cart
	cart,err := carServ.repo.GetCart(id)
	if len(err)>0 {
		return nil,err
	}
	return cart,err
}
func(carServ *CartServ)UpdateCart(id uint,car entity.Cart)(*entity.Cart,[]error){
	cart,ers := carServ.repo.UpdateCart(id,car)

	if len(ers)>0 {
		return nil,ers
	}
	return cart,ers
}
func(carServ *CartServ)CreateCart(car entity.Cart)(*entity.Cart,[]error){

	cart,ers := carServ.repo.CreateCart(car)

	if len(ers)>0 {
		return nil,ers
	}
	return cart,ers
}
func(carServ *CartServ)DeleteCart(id uint)(*entity.Cart,[]error){
	
	cart,ers := carServ.repo.DeleteCart(id)

	if len(ers)>0 {
		return nil,ers
	}
	return cart,ers
}

