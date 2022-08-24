package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/motikingo/ecommerceRESTAPI-Go/catagory"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	"github.com/motikingo/ecommerceRESTAPI-Go/item"
)

type catagoryHandler struct{
	catSrvc catagory.CatagoryService
	itemSrv item.ItemService
	sessionHa *SessionHandler
}

func NewcatHandler(catSrvc catagory.CatagoryService,sessionHa *SessionHandler)catagoryHandler{

	return catagoryHandler{catSrvc: catSrvc, sessionHa: sessionHa}
}

func(catHandler *catagoryHandler) GetCatagories(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		catagoties []entity.Catagory
	}{
		status:"Unauthorized user",
	}
	sess := catHandler.sessionHa.GetSession(ctx)
	if sess != nil{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	catagories,ers:=catHandler.catSrvc.GetCatagories()

	if ers!=nil{
		response.status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	if len(catagories)<1{
		response.status = "No catagories added yet"
		ctx.IndentedJSON(http.StatusOK,response)
		return
	} 
	response.status = "Successfully retrieved catagories"
	response.catagoties = catagories
	ctx.IndentedJSON(200,response)
	
}

func(catHandler *catagoryHandler) GetCatagory(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		catagory *entity.Catagory
	}{
		status:"Unauthorized user",
	}
	sess := catHandler.sessionHa.GetSession(ctx)
	if sess != nil{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	ids := ctx.Param("id")
	id,_:= strconv.Atoi(ids)
	
	catagory,ers:=catHandler.catSrvc.GetCatagory(uint(id))

	if ers!=nil{
		response.status = "No such catagories"
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	response.status = "successful retrieved catagory"
	response.catagory = catagory
	ctx.JSON(200,response)
	
}
func(catHandler *catagoryHandler) CreateCatagory(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		catagory *entity.Catagory
	}{
		status:"Unauthorized user",
	}
	input := &struct{
		name string
		imageurl string 
		description string
		items_Id []uint
	}{}
	sess := catHandler.sessionHa.GetSession(ctx)
	if sess != nil || sess.Role == "Admin"{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	e := ctx.BindJSON(&input)

	if e!=nil || input.name == "" || input.imageurl== "" || input.description == "" || len(input.items_Id)<1{
		response.status = "Incorrect input"
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return

	}
	if catHandler.catSrvc.IsCatagoryNameExist(input.name){
		response.status = "Catagory name already exist"
		ctx.IndentedJSON(http.StatusOK,response)
		return
	}
	
	catagory := &entity.Catagory{
		Name: input.name,
		Image: input.imageurl,
		Description: input.description,
		Items_Id: input.items_Id,
	} 
	catagory,ers:=catHandler.catSrvc.CreateCatagory(*catagory)

	if len(ers)>0{
		response.status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	response.status = "create successfully"
	response.catagory = catagory
	ctx.IndentedJSON(200,response)
}

func(catHandler *catagoryHandler) UpdateCatagory(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		catagory *entity.Catagory
	}{
		status:"Unauthorized user",
	}
	input := &struct{
		name string
		imageurl string 
		description string
	}{}
	sess := catHandler.sessionHa.GetSession(ctx)
	if sess != nil || sess.Role == "Admin"{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	e := ctx.BindJSON(&input)

	if e!=nil || input.name == "" || input.imageurl== "" || input.description == "" {
		response.status = "Incorrect input"
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return

	}
	id,_ := strconv.Atoi(ctx.Param("id"))
	cata,ers := catHandler.catSrvc.GetCatagory(uint(id))
	if len(ers)>0{
		response.status = "Catagory name already exist"
		ctx.IndentedJSON(http.StatusOK,response)
		return
	}
	if cata.Name != input.name{
		if catHandler.catSrvc.IsCatagoryNameExist(input.name){
			response.status = "Catagory name already exist"
			ctx.IndentedJSON(http.StatusOK,response)
			return
		}
		
	}
	catagory := &entity.Catagory{
		Name: input.name,
		Image: input.imageurl,
		Description: input.description,
	} 
	catagory.ID = uint(id)
	catagory,ers = catHandler.catSrvc.UpdateCatagory(*catagory)

	if len(ers)>0{
		response.status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	response.status = "successfully updated"
	response.catagory = catagory
	ctx.IndentedJSON(200,response)
}

func(catHandler *catagoryHandler)AddItems(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		items []entity.Item
	}{
		status:"Unauthorized user",
	}

	input := &struct{
		catagory_Id uint 
		items_id_to_add []int
	}{}
	sess := catHandler.sessionHa.GetSession(ctx)
	if sess != nil || sess.Role == "Admin"{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	e := ctx.BindJSON(&input)
	if e != nil || string(input.catagory_Id) == "" || len(input.items_id_to_add)<1{
		response.status = "Invalid Input"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return

	}
	cata,ers := catHandler.catSrvc.GetCatagory(input.catagory_Id)
	if len(ers)>0{
		response.status = "No such catagory exist"
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}

	for _,itmId := range input.items_id_to_add{
		item,ers := catHandler.itemSrv.GetItem(uint(itmId))
		if len(ers)>0 || item == nil{
			response.status = "No Item with this Id exist"
			ctx.IndentedJSON(http.StatusNotFound,response)
			return
		}

		for _,item_id := range cata.Items_Id{
			if item_id == item.ID {
				response.status = "Wow Item already exist"
				ctx.IndentedJSON(http.StatusOK,response)
				return
			}
		}

		cata.Items_Id = append(cata.Items_Id,item.ID)

		cata,ers = catHandler.catSrvc.UpdateCatagory(*cata)
		if len(ers)>0{
			cata.Items_Id = cata.Items_Id[:len(cata.Items_Id)-2]
			response.status = "Internal Server Error"
			ctx.IndentedJSON(http.StatusInternalServerError,response)
			return
		}
		response.items = append(response.items,*item)
	}

	response.status = "Item Successfully Added Yo! "
	ctx.IndentedJSON(http.StatusInternalServerError,response)
	
}
func(catHandler *catagoryHandler) DeleteItemFromCatagory(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		items []entity.Item
	}{
		status:"Unauthorized user",
	}

	input := &struct{
		catagory_Id uint 
		items_id_to_add []int
	}{}
	sess := catHandler.sessionHa.GetSession(ctx)
	if sess != nil || sess.Role == "Admin"{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	e := ctx.BindJSON(&input)
	if e != nil || string(input.catagory_Id)== "" || len(input.items_id_to_add)<1{
		response.status = "Invalid Input"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return

	}
	cata,ers := catHandler.catSrvc.GetCatagory(input.catagory_Id)
	if len(ers)>0{
		response.status = "No such catagory exist"
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}

	for _,itmId := range input.items_id_to_add{
		item,ers := catHandler.itemSrv.GetItem(uint(itmId))
		if len(ers)>0 || item == nil{
			response.status = "No Item with this Id exist"
			ctx.IndentedJSON(http.StatusNotFound,response)
			return
		}
		check := false
		for _,item_id := range cata.Items_Id{
			if item_id != item.ID {
				cata.Items_Id = append(cata.Items_Id,item.ID)
			}else{
				check = true
			}
		}
		if check{
			cata,ers = catHandler.catSrvc.UpdateCatagory(*cata)
			if len(ers)>0{
				response.status = "Internal Server Error"
				ctx.IndentedJSON(http.StatusInternalServerError,response)
				return
			}
			response.items = append(response.items,*item)

		}
		
	}
	response.status = "Item Successfully Deleted Yo! "
	ctx.IndentedJSON(http.StatusInternalServerError,response)

}

func(catHandler *catagoryHandler) GetMyItems(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		items []entity.Item
	}{
		status:"Unauthorized user",
	}

	input := &struct{
		catagory_Id uint 
	}{}

	e := ctx.BindJSON(&input)
	if e != nil || string(input.catagory_Id) == "" {
		response.status = "Invalid Input"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return

	}

	cata,ers := catHandler.catSrvc.GetCatagory(input.catagory_Id)
	if len(ers)>0{
		response.status = "No such catagory exist"
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}

	for _,itemId := range cata.Items_Id{
		item,ers := catHandler.itemSrv.GetItem(itemId)
		if len(ers)>0{
			response.status = "No such Item exist"
			ctx.IndentedJSON(http.StatusNotFound,response)
			return
		}

		response.items = append(response.items,*item)

	}

	response.status = "all catagory Item retrieved"
	ctx.IndentedJSON(200,response)

}


func(catHandler *catagoryHandler) DeleteCatagory(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		catagory *entity.Catagory
	}{
		status:"Unauthorized user",
	}
	sess := catHandler.sessionHa.GetSession(ctx)
	if sess != nil || sess.Role == "Admin"{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	catagory_Id,_ := strconv.Atoi(ctx.Param("id")) 
	
	catagory,ers:= catHandler.catSrvc.GetCatagory(uint(catagory_Id))

	
	if ers!=nil || catagory == nil {
		response.status = "Catagory not found"
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}

	catagory,ers =catHandler.catSrvc.DeleteCatagory(uint(catagory_Id))

	if ers!=nil{
		response.status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}

	response.status = "catagory successfully deleted"
	response.catagory  = catagory
	ctx.JSON(200,response)
}
