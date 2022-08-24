package handler

import (
	//"fmt"
	"net/http"
	//"net/http"
	//"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/motikingo/ecommerceRESTAPI-Go/cart"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
)

type cartHandler struct{
	carServc cart.CartService 
	sessionHa   *SessionHandler
}

func NewcartHandler(carServc cart.CartService,sessionHa *SessionHandler) cartHandler{
	return cartHandler{carServc: carServc,sessionHa:sessionHa}
}

func(carHandler *cartHandler)GetCarts(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		cart []entity.Cart
		
	}{
		status:"Unauthorized user",
	}

	sess := carHandler.sessionHa.GetSession(ctx)
	if sess != nil{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	carts,ers:= carHandler.carServc.GetCarts()
	if ers!=nil{
		response.status = "Internal Serever Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	response.status = "Successfully retrieved carts"
	response.cart = carts
	ctx.IndentedJSON(200,response)

}

func(carHandler *cartHandler)GetCart(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		cart *entity.Cart
		
	}{
		status:"Unauthorized user",
	}


	id,_ := strconv.Atoi(ctx.Param("id"))
	sess := carHandler.sessionHa.GetSession(ctx)
	if sess != nil{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	cart,ers:= carHandler.carServc.GetCart(uint(id))
	if ers!=nil{
		response.status = "Internal Serever Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	response.status = "Successfully retrieved cart"
	response.cart = cart
	ctx.IndentedJSON(200,response)
}

func(carHandler *cartHandler)CreateCart(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		cart *entity.Cart
		
	}{
		status:"Unauthorized user",
	}
	
	sess := carHandler.sessionHa.GetSession(ctx)
	if sess != nil{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	cart:= carHandler.carServc.GetCartByUserID(sess.UserId)

	if cart != nil{
		response.status = "cart already added"
		response.cart = cart
		ctx.IndentedJSON(http.StatusOK,response)
		return
	}

	cart = &entity.Cart{
		UserId: sess.UserId,
	}

	response.status = "cart added successfully"
	response.cart = cart 
	ctx.IndentedJSON(http.StatusOK,response)
}

func(carHandler *cartHandler)DeleteCart(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		cart *entity.Cart
		
	}{
		status:"Unauthorized user",
	}
	
	sess := carHandler.sessionHa.GetSession(ctx)
	if sess != nil{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	input := &struct{
		cart_Id uint
	}{}

	e := ctx.BindJSON(input)

	if e!=nil || string(input.cart_Id)==""{
		response.status = "Invalid Input"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}
	cart,ers:= carHandler.carServc.DeleteCart(sess.UserId)
	if ers!=nil{
		response.status = "No cart Found"
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}

	response.status = "card deleted successfully"
	response.cart = cart
	ctx.IndentedJSON(200,response)
	
}

