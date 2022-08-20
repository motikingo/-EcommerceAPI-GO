package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/motikingo/ecommerceRESTAPI-Go/catagory"
	"github.com/motikingo/ecommerceRESTAPI-Go/item"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
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
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		catagoties []entity.Catagory
	}{
		status:"Unauthorized user"
	}
	sess := catHa.sessionHa.GetSession(r)
	if sess != nil{
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}
	catagories,ers:=catHandler.catSrvc.GetCatagories()

	if ers!=nil{
		response.status = "Internal Server Error"
		cxt.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	if len(catagories)<1{
		response.status = "No catagories added yet"
		cxt.IndentedJSON(http.StatusOK,response)
		return
	} 
	response.status = "Successfully retrieved catagories"
	response.catagoties = catagories
	ctx.IndentedJSON(200,response)
	
}

func(catHandler *catagoryHandler) GetCatagory(ctx *gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		catagory *entity.Catagory
	}{
		status:"Unauthorized user"
	}
	sess := catHa.sessionHa.GetSession(r)
	if sess != nil{
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}

	ids := ctx.Param("id")
	id,_:= strconv.Atoi(ids)
	
	catagory,ers:=catHandler.catSrvc.GetCatagory(uint(id))

	if ers!=nil{
		response.status = "No such catagories"
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}
	response.status = "successful retrieved catagory"
	response.catagory = catagory
	ctx.JSON(200,&)
	
}
func(catHandler *catagoryHandler) CreateCatagory(ctx *gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		catagory *entity.Catagory
	}{
		status:"Unauthorized user"
	}
	input := &struct{
		name string
		imageurl string 
		description string
		items_Id []uint
	}{}
	sess := catHa.sessionHa.GetSession(r)
	if sess != nil || sess.Role == "Admin"{
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}
	e := ctx.BindJSON(&input)

	if e!=nil || input.name == "" || input.imageurl== "" || input.description == "" || len(input_Id)<1{
		response.status = "Incorrect input"
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return

	}
	catagories,ers:=catHandler.catSrvc.Catagories()
	if len(ers)>0{
		response.status = "Internal Server Error"
		cxt.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	if len(catagories)<1{
		response.status = "No catagories exist"
		cxt.IndentedJSON(http.StatusOK,response)
		return
	}
	for _,cata:= range catagories{
		if cata.Name == input.name{
			response.status = "Catagory name already exist"
			cxt.IndentedJSON(http.StatusOK,response)
			return
		}
	}
	catagory := &entity.Catagory{
		Name: input.name
		Image: input.imageurl
		Description: input.description
		Items_Id: input.items_Id
	} 
	catagory,ers:=catHandler.catSrvc.CreateCatagory(catagory)

	if len(ers)>0{
		response.status = "Internal Server Error"
		cxt.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	response.status = "create successfully"
	response.catagory = catagory
	ctx.IndentedJSON(200,response)
}

func(catHandler *catagoryHandler) UpdateCatagory(ctx *gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		catagory *entity.Catagory
	}{
		status:"Unauthorized user"
	}
	input := &struct{
		name string
		imageurl string 
		description string
	}{}
	sess := catHa.sessionHa.GetSession(r)
	if sess != nil || sess.Role == "Admin"{
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}
	e := ctx.BindJSON(&input)

	if e!=nil || input.name == "" || input.imageurl== "" || input.description == "" {
		response.status = "Incorrect input"
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return

	}
	catagories,ers:=catHandler.catSrvc.Catagories()
	if len(ers)>0{
		response.status = "Internal Server Error"
		cxt.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	if len(catagories)<1{
		response.status = "No catagories exist"
		cxt.IndentedJSON(http.StatusOK,response)
		return
	}
	var catagory_Id string

	id,_:= strconv.Atoi(ctx.BindJson(catagory_Id))

	for _,cata:= range catagories{
		if cata.Name == input.name && cata.ID != uint(id){
			response.status = "Catagory name already exist"
			cxt.IndentedJSON(http.StatusOK,response)
			return
		}
	}
	catagory := &entity.Catagory{
		Name: input.name
		Image: input.imageurl
		Description: input.description
	} 
	catagory.ID = uint(id)
	catagory,ers:=catHandler.catSrvc.UpdateCatagory(catagory)

	if len(ers)>0{
		response.status = "Internal Server Error"
		cxt.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	resporesponse.status = "successfully updated"
	response.catagory = input.catagory
	ctx.IndentedJSON(200,response)
}

func(catHandler *catagoryHandler)AddItems(ctx *gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		items []entity.Item
	}{
		status:"Unauthorized user"
	}

	input := &struct{
		catagory_Id uint 
		items_id_to_add []int
	}{}
	sess := catHa.sessionHa.GetSession(r)
	if sess != nil || sess.Role == "Admin"{
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}
	e := ctx.BindJSON(&input)
	if e != nil || input.catagory_Id == nil || len(input.items_id_to_add)<1{
		response.status = "Invalid Input"
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return

	}
	cata,ers := catHandler.catSrv.GetCatagory(input.catagory_Id)
	if len(ers)>0{
		response.status = "No such catagory exist"
		cxt.IndentedJSON(http.StatusNotFound,response)
		return
	}

	for _,itmId := range input.item_Id_to_add{
		item,ers := catHandler.itemSrv.Item(uint(itmId))
		if len(ers)>0 || item == nil{
			response.status = "No Item with this Id exist"
			cxt.IndentedJSON(http.StatusNotFound,response)
			return
		}

		for _,item_id := cata.Items_Id{
			if item_id == item.ID {
				response.status = "Wow Item already exist"
				cxt.IndentedJSON(http.StatusOK,response)
				return
			}
		}

		cata.Item_Id = append(cata.Item_Id,item.ID)

		cata,ers := catcatHandler.catSrvc.UpdateCatagory(cata)
		if len(ers)>0{
			cata.Item_Id = cat_Item_Id[:len(cata.Item_Id)-2]
			response.status = "Internal Server Error"
			cxt.IndentedJSON(http.StatusInternalServerError,response)
			return
		}
		response.items = append(response.items,*item)
	}

	response.status = "Item Successfully Added Yo! "
	cxt.IndentedJSON(http.StatusInternalServerError,response)
	
}
func(catHandler *catagoryHandler) DeleteItemFromCatagory(ctx *gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		items []entity.Item
	}{
		status:"Unauthorized user"
	}

	input := &struct{
		catagory_Id uint 
		items_id_to_add []int
	}{}
	sess := catHa.sessionHa.GetSession(r)
	if sess != nil || sess.Role == "Admin"{
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}
	e := ctx.BindJSON(&input)
	if e != nil || input.catagory_Id == nil || len(input.items_id_to_add)<1{
		response.status = "Invalid Input"
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return

	}
	cata,ers := catHandler.catSrv.GetCatagory(input.catagory_Id)
	if len(ers)>0{
		response.status = "No such catagory exist"
		cxt.IndentedJSON(http.StatusNotFound,response)
		return
	}

	for _,itmId := range input.item_Id_to_add{
		item,ers := catHandler.itemSrv.Item(uint(itmId))
		if len(ers)>0 || item == nil{
			response.status = "No Item with this Id exist"
			cxt.IndentedJSON(http.StatusNotFound,response)
			return
		}
		check := false
		for _,item_id := cata.Items_Id{
			if item_id != item.ID {
				cata.Item_Id = append(cata.Item_Id,item.ID)
			}else{
				check = true
			}
		}
		if check{
			cata,ers := catcatHandler.catSrvc.UpdateCatagory(cata)
			if len(ers)>0{
				response.status = "Internal Server Error"
				cxt.IndentedJSON(http.StatusInternalServerError,response)
				return
			}
			response.items = append(response.items,*item)

		}
		
	}
	response.status = "Item Successfully Deleted Yo! "
	cxt.IndentedJSON(http.StatusInternalServerError,response)

}

func(catHandler *catagoryHandler) GetMyItems(ctx *gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		items []entity.Item
	}{
		status:"Unauthorized user"
	}

	input := &struct{
		catagory_Id uint 
	}{}

	e := ctx.BindJSON(&input)
	if e != nil || input.catagory_Id == nil {
		response.status = "Invalid Input"
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return

	}

	cata,ers := catHandler.catSrv.GetCatagory(input.catagory_Id)
	if len(ers)>0{
		response.status = "No such catagory exist"
		cxt.IndentedJSON(http.StatusNotFound,response)
		return
	}

	for _,itemId := range cata.Item_Id{
		item,ers := catcatHandler.itemSrv.Item(itemId)
		if len(ers)>0{
			response.status = "No such Item exist"
			cxt.IndentedJSON(http.StatusNotFound,response)
			return
		}

		response.items = append(response.items,*item)

	}

	response.status := "all catagory Item retrieved"
	ctx.IndentedJSON(200,response)

}


func(catHandler *catagoryHandler) DeleteCatagory(ctx *gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		catagory *entity.Catagory
	}{
		status:"Unauthorized user"
	}
	sess := catHa.sessionHa.GetSession(ctx)
	if sess != nil || sess.Role == "Admin"{
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}

	var catagory_Id string
	id,e:= strconv.Atoi(ctx.BindJSON(catagory_Id))
	if e!=nil{
		response.status = "Json format Incorrect"
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	catagory,ers:= catHandler.catSrvc.Catagory(uint(id))

	
	if ers!=nil || catagory == nil {
		response.status = "Catagory not found"
		cxt.IndentedJSON(http.StatusNotFound,response)
		return
	}

	catagory,ers =catHandler.catSrvc.DeleteCatagory(uint(id))

	if ers!=nil{
		response.status = "Internal Server Error"
		cxt.IndentedJSON(http.StatusInternalServerError,response)
		return
	}

	response.status = "catagory successfully deleted"
	resporesponse.catagory  = catagory
	ctx.JSON(200,response)
}
