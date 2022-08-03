package handler

import (
	//"fmt"
	"log"
	//"net/http"
	//"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/motikingo/ecommerceRESTAPI-Go/cart"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
)

type cartHandler struct{
	carServc cart.CartService 
}

func NewcartHandler(carServc cart.CartService) cartHandler{
	return cartHandler{carServc: carServc}
}

func(carHandler *cartHandler)GetCarts(ctx *gin.Context){

	carts,ers:= carHandler.carServc.GetCarts()
	if ers!=nil{
		log.Fatal(ers)
	}

	ctx.JSON(200,&carts)

}

func(carHandler *cartHandler)GetCart(ctx *gin.Context){
	
	ids := ctx.Param("id")
	id,er := strconv.Atoi(ids)
	if er!=nil{
		log.Fatal(er)
	}
	cart,ers:= carHandler.carServc.GetCart(uint(id))
	if ers!=nil{
		log.Fatal(ers)
	}

	ctx.JSON(200,&cart)
}

func(carHandler *cartHandler)UpdateCart(ctx *gin.Context){
	var cart entity.Cart
	ids:= ctx.Param("id")

	id,e:= strconv.Atoi(ids)
	if e!=nil{
		log.Fatal(e)
	}
	er := ctx.BindJSON(&cart)
	
	if er!=nil{
		log.Fatal(er)
	}
	car,ers:= carHandler.carServc.UpdateCart(uint(id), cart)
	if ers!=nil{
		log.Fatal(ers)
	}

	ctx.JSON(200,&car)
}

func(carHandler *cartHandler)CreateCarts(ctx *gin.Context){

	var cart entity.Cart
	er := ctx.BindJSON(&cart)
	
	if er!=nil{
		log.Fatal(er)
	}
	car,ers:= carHandler.carServc.CreateCart(cart)
	if ers!=nil{
		log.Fatal(ers)
	}

	ctx.JSON(200,&car)
	
}

func(carHandler *cartHandler)DeleteCarts(ctx *gin.Context){
	ids := ctx.Param("id")
	id,er := strconv.Atoi(ids)
	if er!=nil{
		log.Fatal(er)
	}
	cart,ers:= carHandler.carServc.DeleteCart(uint(id))
	if ers!=nil{
		log.Fatal(ers)
	}

	ctx.JSON(200,&cart)
	
}

