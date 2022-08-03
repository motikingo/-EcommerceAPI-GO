package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/motikingo/ecommerceRESTAPI-Go/cart"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
)

type CartRepo struct{
	db *gorm.DB
}

func NewCartRepo(db *gorm.DB) cart.CartRepository{
	return &CartRepo{db: db}
}

func(carRepo *CartRepo) GetCarts()([]entity.Cart,[]error){

	var carts []entity.Cart
	err := carRepo.db.Find(&carts).GetErrors()
	if len(err)>0 {
		return nil,err
	}
	return carts,err

}
func(carRepo *CartRepo)GetCart(id uint)(*entity.Cart,[]error){

	var cart entity.Cart
	err := carRepo.db.First(&cart).GetErrors()
	if len(err)>0 {
		return nil,err
	}
	return &cart,err
}
func(carRepo *CartRepo)UpdateCart(id uint,car entity.Cart)(*entity.Cart,[]error){

	cart,ers:= carRepo.GetCart(id)

	if len(ers)>0 {
		return nil,ers
	}
	cart.Items = car.Items
	ers = carRepo.db.Save(&cart).GetErrors()

	if len(ers)>0 {
		return nil,ers
	}
	return cart,ers
}
func(carRepo *CartRepo)CreateCart(car entity.Cart)(*entity.Cart,[]error){

	// if len(ers)>0 {
	// 	return nil,ers
	// }
	cart := car
	ers := carRepo.db.Create(&cart).GetErrors()

	if len(ers)>0 {
		return nil,ers
	}
	return &cart,ers
}
func(carRepo *CartRepo)DeleteCart(id uint)(*entity.Cart,[]error){
	cart,ers:= carRepo.GetCart(id)

	ers = carRepo.db.Delete(&cart).GetErrors()

	if len(ers)>0 {
		return nil,ers
	}
	return cart,ers
}



