package handler

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/motikingo/ecommerceRESTAPI-Go/catagory"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
)

type catagoryHandler struct{
	catSrvc catagory.CatagoryService
}

func NewcatHandler(catSrvc catagory.CatagoryService)catagoryHandler{

	return catagoryHandler{catSrvc: catSrvc}
}

func(catHandler *catagoryHandler) GetCatagories(ctx *gin.Context){
	catagories,ers:=catHandler.catSrvc.GetCatagories()

	if ers!=nil{
		log.Fatal(ers)
	}
	ctx.JSON(200,&catagories)
	
}

func(catHandler *catagoryHandler) GetCatagory(ctx *gin.Context){
	ids := ctx.Param("id")
	id,e:= strconv.Atoi(ids)
	if e!=nil{
		log.Fatal(e)
	}
	catagories,ers:=catHandler.catSrvc.GetCatagory(uint(id))

	if ers!=nil{
		log.Fatal(ers)
	}
	ctx.JSON(200,&catagories)
	
}

func(catHandler *catagoryHandler) UpdateCatagory(ctx *gin.Context){
	var cat entity.Catagory 
	ids := ctx.Param("id")
	id,e:= strconv.Atoi(ids)
	if e!=nil{
		log.Fatal(e)
	}

	e = ctx.BindJSON(&cat)

	if e!=nil{
		log.Fatal(e)
	}

	catagory,ers:=catHandler.catSrvc.UpdateCatagory(uint(id),cat)

	if ers!=nil{
		log.Fatal(ers)
	}
	ctx.JSON(200,&catagory)
}

func(catHandler *catagoryHandler) CreateCatagory(ctx *gin.Context){
	var cat entity.Catagory 
	e := ctx.BindJSON(&cat)

	if e!=nil{
		log.Fatal(e)
	}

	catagory,ers:=catHandler.catSrvc.CreateCatagory(cat)

	if ers!=nil{
		log.Fatal(ers)
	}
	ctx.JSON(200,&catagory)
	
}

func(catHandler *catagoryHandler) DeleteCatagory(ctx *gin.Context){
	ids := ctx.Param("id")
	id,e:= strconv.Atoi(ids)
	if e!=nil{
		log.Fatal(e)
	}

	catagory,ers:=catHandler.catSrvc.DeleteCatagory(uint(id))

	if ers!=nil{
		log.Fatal(ers)
	}
	ctx.JSON(200,&catagory)
}
