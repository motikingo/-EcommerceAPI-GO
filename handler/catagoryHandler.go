package handler

import (
	"fmt"
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

func NewcatHandler(catSrvc catagory.CatagoryService,itemSrv item.ItemService,sessionHa *SessionHandler)catagoryHandler{

	return catagoryHandler{catSrvc: catSrvc,itemSrv:itemSrv, sessionHa: sessionHa}
}

func(catHandler *catagoryHandler) GetCatagories(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		Status string
		Catagoties []entity.Catagory
	}{
		Status:"Unauthorized user",
	}
	sess := catHandler.sessionHa.GetSession(ctx)
	if sess == nil{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	catagories,ers:=catHandler.catSrvc.GetCatagories()

	if len(ers)>0{
		response.Status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	if len(catagories)<1{
		response.Status = "No catagories added yet"
		ctx.IndentedJSON(http.StatusOK,response)
		return
	} 
	response.Status = "Successfully retrieved catagories"
	response.Catagoties = catagories
	ctx.IndentedJSON(200,response)
	
}

func(catHandler *catagoryHandler) GetCatagory(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		Status string
		Catagory *entity.Catagory
	}{
		Status:"Unauthorized user",
	}
	sess := catHandler.sessionHa.GetSession(ctx)
	if sess == nil{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	ids := ctx.Param("id")
	id,_:= strconv.Atoi(ids)
	
	catagory,ers:=catHandler.catSrvc.GetCatagory(uint(id))

	if ers!=nil{
		response.Status = "No such catagories"
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	response.Status = "successful retrieved catagory"
	response.Catagory = catagory
	ctx.JSON(200,response)
	
}
func(catHandler *catagoryHandler) CreateCatagory(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		Status string
		Catagory *entity.Catagory
	}{
		Status:"Unauthorized user",
	}
	input := &struct{
		Name string
		Description string
		Imageurl string 
		Items_Id []uint
	}{}
	sess := catHandler.sessionHa.GetSession(ctx)
	if sess == nil || sess.Role != "Admin"{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	e := ctx.BindJSON(&input)

	if e!=nil || input.Name == "" || input.Imageurl== "" || input.Description == "" || len(input.Items_Id)<1{
		response.Status = "Incorrect input"
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return

	}
	if catHandler.catSrvc.IsCatagoryNameExist(input.Name){
		response.Status = "Catagory name already exist"
		ctx.IndentedJSON(http.StatusOK,response)
		return
	}
	
	catagory := &entity.Catagory{
		Name: input.Name,
		Description: input.Description,
		Image: input.Imageurl,

	} 
	for _,id := range input.Items_Id{
		item,er := catHandler.itemSrv.GetItem(id)
		if len(er)>0{
			response.Status = "No Item found"
			ctx.IndentedJSON(http.StatusInternalServerError,response)
			return
		}
		catagory.Items = append(catagory.Items, *item)
	}
	catagory,ers:=catHandler.catSrvc.CreateCatagory(*catagory)

	if len(ers)>0{
		response.Status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	response.Status = "create successfully"
	response.Catagory = catagory
	ctx.IndentedJSON(200,response)
}

func(catHandler *catagoryHandler) UpdateCatagory(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		Status string
		Catagory *entity.Catagory
	}{
		Status:"Unauthorized user",
	}
	input := &struct{
		Name string
		Imageurl string 
		Description string
	}{}
	sess := catHandler.sessionHa.GetSession(ctx)
	if sess == nil || sess.Role != "Admin"{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	e := ctx.BindJSON(&input)

	if e!=nil || input.Name == "" || input.Imageurl== "" || input.Description == "" {
		response.Status = "Incorrect input"
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return

	}
	id,_ := strconv.Atoi(ctx.Param("id"))
	cata,ers := catHandler.catSrvc.GetCatagory(uint(id))
	if len(ers)>0{
		response.Status = "No such Catagory"
		ctx.IndentedJSON(http.StatusOK,response)
		return
	}
	if cata.Name != input.Name{
		if catHandler.catSrvc.IsCatagoryNameExist(input.Name){
			response.Status = "Catagory name already exist"
			ctx.IndentedJSON(http.StatusOK,response)
			return
		}
		
	}
	
	cata.Name = input.Name
	cata.Image = input.Imageurl
	cata.Description = input.Description
	catagory,ers := catHandler.catSrvc.UpdateCatagory(*cata)

	if len(ers)>0{
		response.Status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	response.Status = "successfully updated"
	response.Catagory = catagory
	ctx.IndentedJSON(200,response)
}

func(catHandler *catagoryHandler)AddItems(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		Status string
		Catagory entity.Catagory
	}{
		Status:"Unauthorized user",
	}

	input := &struct{
		Items_id_to_add []uint
	}{}
	sess := catHandler.sessionHa.GetSession(ctx)
	if sess == nil || sess.Role != "Admin"{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	
	if e := ctx.BindJSON(&input); e != nil || len(input.Items_id_to_add)<1{
		response.Status = "Invalid Input"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return

	}
	id,_ := strconv.Atoi(ctx.Param("id"))
	cata,ers := catHandler.catSrvc.GetCatagory(uint(id))
	if len(ers)>0{
		response.Status = fmt.Sprintln("No catagory exist with id: %d",id)
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}

	for _,itmId := range input.Items_id_to_add{
		item,ers := catHandler.itemSrv.GetItem(uint(itmId))
		if len(ers)>0 || item == nil{
			response.Status = fmt.Sprintln("No Item with this Id %d exist",itmId)
			ctx.IndentedJSON(http.StatusNotFound,response)
			return
		}

		for _,itm := range cata.Items{
			if itm.ID == item.ID {
				response.Status = "Wow Item already exist"
				ctx.IndentedJSON(http.StatusOK,response)
				return
			}
		}

		cata.Items = append(cata.Items,*item)
		
	}
	_,ers = catHandler.catSrvc.UpdateCatagory(*cata)

	if len(ers)>0{
		response.Status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	response.Status = "Item Successfully Added Yo! "
	response.Catagory = *cata
	ctx.IndentedJSON(http.StatusInternalServerError,response)
	
}

func(catHandler *catagoryHandler) DeleteItemFromCatagory(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		Status string
		Catagory entity.Catagory
	}{
		Status:"Unauthorized user",
	}

	input := &struct{
		Items_id_to_add []int
	}{}
	sess := catHandler.sessionHa.GetSession(ctx)
	if sess == nil || sess.Role != "Admin"{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	e := ctx.BindJSON(&input)
	if e != nil || len(input.Items_id_to_add)<1{
		response.Status = "Invalid Input"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return

	}
	id,_ := strconv.Atoi(ctx.Param("id"))
	cata,ers := catHandler.catSrvc.GetCatagory(uint(id))
	if len(ers)>0{
		response.Status = fmt.Sprintf("No catagory with id %d exist",id)
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}

	for _,itmId := range input.Items_id_to_add{
		item,ers := catHandler.itemSrv.GetItem(uint(itmId))
		if len(ers)>0 || item == nil{
			response.Status = fmt.Sprintf("No Item with this Id %d exist",item.ID)
			ctx.IndentedJSON(http.StatusNotFound,response)
			return
		}
		check := false
		var items []entity.Item
		for _,itm := range cata.Items{
			if itm.ID != item.ID {
				items = append(items,*item)
			}else{
				fmt.Println("here")
				check = true
				break
			}
		}
		if check{
			cata.Items = items
			cata,ers = catHandler.catSrvc.UpdateCatagory(*cata)
			if len(ers)>0{
				response.Status = "Internal Server Error"
				ctx.IndentedJSON(http.StatusInternalServerError,response)
				return
			}
		}
		
	}
	response.Status = "Item Successfully Deleted Yo! "
	response.Catagory = *cata
	ctx.IndentedJSON(http.StatusOK,response)
	

}

func(catHandler *catagoryHandler) GetMyItems(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		Status string
		Items []entity.Item
	}{
		Status:"Unauthorized user",
	}

	input := &struct{
		Catagory_Id uint 
	}{}

	e := ctx.BindJSON(&input)
	if e != nil || string(input.Catagory_Id) == "" {
		response.Status = "Invalid Input"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return

	}
	cata,ers := catHandler.catSrvc.GetCatagory(input.Catagory_Id)
	if len(ers)>0{
		response.Status = "No such catagory exist"
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}

	for _,itm := range cata.Items{
		item,ers := catHandler.itemSrv.GetItem(itm.ID)
		if len(ers)>0{
			response.Status = "No such Item exist"
			ctx.IndentedJSON(http.StatusNotFound,response)
			return
		}

		response.Items = append(response.Items,*item)

	}

	response.Status = "all catagory Item retrieved"
	ctx.IndentedJSON(200,response)

}


func(catHandler *catagoryHandler) DeleteCatagory(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		Status string
		Catagory *entity.Catagory
	}{
		Status:"Unauthorized user",
	}
	sess := catHandler.sessionHa.GetSession(ctx)
	if sess == nil || sess.Role != "Admin"{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	catagory_Id,_ := strconv.Atoi(ctx.Param("id")) 
	
	catagory,ers:= catHandler.catSrvc.GetCatagory(uint(catagory_Id))

	
	if ers!=nil || catagory == nil {
		response.Status = "Catagory not found"
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}

	catagory,ers =catHandler.catSrvc.DeleteCatagory(uint(catagory_Id))

	if ers!=nil{
		response.Status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}

	response.Status = "catagory successfully deleted"
	response.Catagory  = catagory
	ctx.JSON(200,response)
}
