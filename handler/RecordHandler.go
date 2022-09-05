package handler

import (
	//"fmt"
	"net/http"
	//"net/http"
	//"encoding/json"
	//"strconv"

	"github.com/gin-gonic/gin"
	"github.com/motikingo/ecommerceRESTAPI-Go/record"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
)

type RecordHandler struct{
	recoSrv record.RecordService
	cartHa 		CartHandler
	sessionHa   *SessionHandler
}

func NewRecordHandler(recoSrv record.RecordService,cartHa CartHandler,sessionHa*SessionHandler) RecordHandler{
	return RecordHandler{recoSrv: recoSrv,cartHa:cartHa,sessionHa:sessionHa}
}

func(recoHandler *RecordHandler)GetRecords(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		Status string
		Record []entity.Record
		
	}{
		Status:"Unauthorized user",
	}

	sess := recoHandler.sessionHa.GetSession(ctx)
	if sess != nil{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	records:= recoHandler.recoSrv.GetRecords()
	if len(records)<1{
		response.Status = "Internal Serever Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	response.Status = "Successfully retrieved records"
	response.Record = records
	ctx.IndentedJSON(200,response)

}

func(recoHandler *RecordHandler)GetRecord(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		Status string
		Record *entity.Record
		
	}{
		Status:"Unauthorized user",
	}


	//id,_ := strconv.Atoi(ctx.Param("id"))
	sess := recoHandler.sessionHa.GetSession(ctx)
	if sess == nil{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	record:= recoHandler.recoSrv.GetRecordByUserID(sess.UserId)
	if record ==nil{
		response.Status = "You Don't have any record yet"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	response.Status = "Successfully retrieved cart"
	response.Record = record
	ctx.IndentedJSON(200,response)
}

// func(recoHandler *RecordHandler)CreateCart(ctx *gin.Context){
// 	ctx.Header("Content-Type","application/json")
// 	response := &struct{
// 		status string
// 		cart *entity.Cart
		
// 	}{
// 		status:"Unauthorized user",
// 	}
	
// 	sess := recoHandler.sessionHa.GetSession(ctx)
// 	if sess != nil{
// 		ctx.IndentedJSON(http.StatusUnauthorized,response)
// 		return
// 	}
// 	cart:= recoHandler.carServc.GetCartByUserID(sess.UserId)

// 	if cart != nil{
// 		response.status = "cart already added"
// 		response.cart = cart
// 		ctx.IndentedJSON(http.StatusOK,response)
// 		return
// 	}

// 	cart = &entity.Cart{
// 		UserId: sess.UserId,
// 	}

// 	response.status = "cart added successfully"
// 	response.cart = cart 
// 	ctx.IndentedJSON(http.StatusOK,response)
// }

func(recoHandler *RecordHandler)ClearRecord(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		Status string
		Record *entity.Record
		
	}{
		Status:"Unauthorized user",
	}
	
	sess := recoHandler.sessionHa.GetSession(ctx)
	if sess != nil{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	input := &struct{
		Record_Id uint
	}{}

	

	if e := ctx.BindJSON(input); e!=nil || string(input.Record_Id)==""{
		response.Status = "Invalid Input"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}
	record:= recoHandler.recoSrv.ClearRecord(sess.UserId)
	if record!=nil{
		response.Status = "No reoord Found"
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}

	response.Status = "card deleted successfully"
	response.Record = record
	ctx.IndentedJSON(200,response)
	
}

