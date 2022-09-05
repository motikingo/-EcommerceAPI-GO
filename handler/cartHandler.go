package handler

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/motikingo/ecommerceRESTAPI-Go/record"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
)
const(
	cartSessionName = "cart" 
)
var cartKey = []byte("super cart")

type CartHandler struct {
	recordSrv record.RecordService
	sessionHa *SessionHandler
}

func NewcartHandler(recordSrv record.RecordService,sessionHa *SessionHandler)CartHandler  {
	return CartHandler{recordSrv:recordSrv,sessionHa:sessionHa}
}

func(car *CartHandler) CreateCart(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response :=&struct{
		Message string
	}{
		Message:"UnAthorized User",
	}	
	sess:= car.sessionHa.GetSession(ctx)
	if sess == nil{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	cart := &entity.Cart{
		UserId:sess.UserId,
		Items:make(map[uint]entity.ItemInfo),
	}
	expireTime := time.Now().Add(2*time.Hour)

	cart.StandardClaims = jwt.StandardClaims{
		ExpiresAt:expireTime.Unix(),
	}
	
	tkn:= jwt.NewWithClaims(jwt.SigningMethodHS256,cart)
	tknstr,er := tkn.SignedString(cartKey)

	if er != nil {
		response.Message = "Internal  server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	cookie := & http.Cookie{
		Name:cartSessionName,
		Value:tknstr,
		Expires:expireTime,
		Path:"/",
		HttpOnly:true,

	}
	http.SetCookie(ctx.Writer,cookie)
	response.Message = "Cart Added"
	ctx.IndentedJSON(http.StatusOK,response)

}

func(car *CartHandler) UpdateCart(cart entity.Cart,ctx *gin.Context)bool{
	ctx.Header("Content-Type","application/json")
	response :=&struct{
		Message string
	}{
		Message:"UnAthorized User",
	}	
	sess:= car.sessionHa.GetSession(ctx)
	if sess == nil{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return false
	}
	expireTime := time.Now().Add(2 * time.Hour)
	cart.StandardClaims = jwt.StandardClaims{
		ExpiresAt:expireTime.Unix(),
	}
	
	tkn:= jwt.NewWithClaims(jwt.SigningMethodHS256,cart)
	tknstr,er := tkn.SignedString(cartKey)

	if er != nil {
		response.Message = "Internal  server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return false
	}
	cookie := http.Cookie{
		Name: cartSessionName,
		Value: tknstr,
		Path: "/",
		Expires: expireTime,
		HttpOnly: true,
	}
	http.SetCookie(ctx.Writer,&cookie)
	//ctx.SetCookie(cartSessionName,tknstr,int(expireTime.Unix()),"/","",true,true)
	response.Message = "Cart Added"
	ctx.IndentedJSON(http.StatusOK,response)
	return true
}
func(car *CartHandler) GetCart(ctx *gin.Context) *entity.Cart{
	ctx.Header("Content-Type","application/json")
	response := &struct{
		Message string
		Cart entity.Cart
	}{
		Message:"Unauthorized User",
	}
	sess:= car.sessionHa.GetSession(ctx)
	if sess == nil{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return nil
	}

	cookie,er := ctx.Request.Cookie(cartSessionName)

	if er != nil || cookie == nil{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return nil
	}
	var cart entity.Cart 
	tknstr,er := jwt.ParseWithClaims(cookie.Value,&cart,func (token *jwt.Token)(interface {},error){
		return cartKey,nil	
	})

	if er !=nil || !tknstr.Valid{
		response.Message = "Invalid Session"
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return nil
	}

	return &cart
}

func(car *CartHandler) DeleteCart(ctx *gin.Context)*entity.Cart{

	ctx.Header("Content-Type","application/json")
	response := &struct{
		Message string
	}{
		Message:"Unauthorized User",
	}
	sess:= car.sessionHa.GetSession(ctx)
	if sess == nil{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return nil
	}

	cart := car.GetCart(ctx)
	if cart == nil{
		response.Message = "No cart found"
		ctx.IndentedJSON(http.StatusNotFound,response)
		return nil
	}

	var c entity.Cart
	expireTime := time.Now().Add(-1*time.Hour)
	c.StandardClaims = jwt.StandardClaims{
		ExpiresAt:expireTime.Unix(),
	}
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256,c)
	tknstr,er := tkn.SignedString(cartKey) 
	if er != nil{
		response.Message = "Internal  server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return nil
	}
	cookie := http.Cookie{
		Name: cartSessionName,
		Value: tknstr,
		Path: "/",
		Expires: expireTime,
		HttpOnly: true,
	}
	http.SetCookie(ctx.Writer,&cookie)
	//ctx.SetCookie(cartSessionName,tknstr,int(expireTime.Unix()),"/","",true,true)
	response.Message = "Cart Deleted Successfully"
	ctx.IndentedJSON(http.StatusOK,response)
	return cart
}
