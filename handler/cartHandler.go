package handler

import (
	//"fmt"
	"log"
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
}

func NewcartHandler(carServc cart.CartService) cartHandler{
	return cartHandler{carServc: carServc}
}

func(carHandler *cartHandler)GetCarts(ctx *gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		cart []entity.Cart
		
	}{
		status:"Unauthorized user"
	}

	sess := catHa.sessionHa.GetSession(ctx)
	if sess != nil{
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}
	carts,ers:= carHandler.carServc.GetCarts()
	if ers!=nil{
		response.status = "Internal Serever Error"
		cxt.IndentedJSON(http.StatusInternalSereverError,response)
		return
	}
	response.status = "Successfully retrieved carts"
	response.cart = carts
	ctx.IndentedJSON(200,response)

}

func(carHandler *cartHandler)GetCart(ctx *gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		cart *entity.Cart
		
	}{
		status:"Unauthorized user"
	}


	id := strconv.Atoi(ctx.Param("id"))
	sess := catHa.sessionHa.GetSession(ctx)
	if sess != nil{
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}

	cart,ers:= carHandler.carServc.GetCart(uint(id))
	if ers!=nil{
		response.status = "Internal Serever Error"
		cxt.IndentedJSON(http.StatusInternalSereverError,response)
		return
	}
	response.status = "Successfully retrieved cart"
	response.cart = cart
	ctx.IndentedJSON(200,response)
}

func(carHandler *cartHandler)CreateCart(ctx *gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		cart *entity.Cart
		
	}{
		status:"Unauthorized user"
	}
	
	sess := catHa.sessionHa.GetSession(ctx)
	if sess != nil{
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}
	cart:= carHandler.carServc.GetCartByUserID(sess.User_Id)

	if cart != nil{
		response.status = "cart already added"
		response.cart = cart
		cxt.IndentedJSON(http.StatusOK,response)
		return
	}

	cart = entity.Cart{
		User_Id: sess.User_Id
	}

	response.status = "cart added successfully"
	response.cart = append(response.cart,cart) 
	cxt.IndentedJSON(http.StatusOK,response)
}

func(carHandler *cartHandler)DeleteCarts(ctx *gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		cart *entity.Cart
		
	}{
		status:"Unauthorized user"
	}
	
	sess := catHa.sessionHa.GetSession(ctx)
	if sess != nil{
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}
	input := &struct{
		cart_Id uint
	}{}

	e := ctx.BindJSON(input)

	if er!=nil || input.cart_Id==nil{
		response.status = "Invalid Input"
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}
	cart,ers:= carHandler.carServc.DeleteCart(uint(id))
	if ers!=nil{
		response.status = "No cart Found"
		cxt.IndentedJSON(http.StatusNotFound,response)
		return
	}

	response.status = "card deleted successfully"
	response.cart = cart
	ctx.IndentedJSON(200,response)
	
}

