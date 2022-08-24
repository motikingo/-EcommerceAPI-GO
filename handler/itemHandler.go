package handler

import (
	//"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/motikingo/ecommerceRESTAPI-Go/catagory"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	"github.com/motikingo/ecommerceRESTAPI-Go/item"
)

type ItemHandler struct{
	itemServ item.ItemService
	sessionHa *SessionHandler
	cataSrv catagory.CatagoryService
}

func NewItemHandler(itemServ item.ItemService, sessionHa *SessionHandler)ItemHandler{
	return ItemHandler{itemServ: itemServ,sessionHa:sessionHa}
}

func (itemHandler *ItemHandler)GetItems(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")

	response := &struct{
		status string
		items [] entity.Item
	}{
		status:"Unauthorized user",
	}

	sess := itemHandler.sessionHa.GetSession(ctx)
	if sess != nil {
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	items,ers:=itemHandler.itemServ.GetItems()
	if  len (ers)>0 {
		response.status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}

	if len(items)==0 {
		response.status = "No item added yet"
		ctx.IndentedJSON(http.StatusOK,response)
		return
	}
	response.status = "items retrieved successfully"
	response.items = items
	ctx.JSON(200,items)


}
func (itemHandler *ItemHandler)GetItem(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		item *entity.Item
	}{
		status:"Unauthorized user",
	}

	sess := itemHandler.sessionHa.GetSession(ctx)
	if sess != nil {
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}


	id,e:= strconv.Atoi(ctx.Param("id")) 
	if  e != nil {
		response.status = "bad request..."
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	item,ers:=itemHandler.itemServ.GetItem(uint(id))
	if  len (ers)>0 {
		response.status = "No such Item"
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}
	response.status = "Item fetched"
	response.item = item
	ctx.IndentedJSON(http.StatusOK,response)
}

func (itemHandler *ItemHandler)CreateItem(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		item *entity.Item
	}{
		status:"Unauthorized user",
	}
	input := &struct{
		name string
		description string
		brand string
		imageurl string
		price float64
		number int
		production_date time.Time
		expire_date time.Time
	}{}
	sess := itemHandler.sessionHa.GetSession(ctx)
	if sess != nil || sess.Role !="Admin"{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	e := ctx.BindJSON(&input)
	if e != nil{
		response.status = "failed to bind input"
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}
	if input.name == "" || input.description== "" || input.brand == "" || input.imageurl == "" || input.production_date.String() == "" || input.expire_date.String()== "" || input.price <= 0 || input.number <=0{
		response.status = "Incorrect input..."
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}

	if input.expire_date.Before(time.Now()){
		response.status = "this Item is expired..."
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}
	items,ers := itemHandler.itemServ.GetItems()

	if len(ers)>0{
		response.status = "Internal Server Error ..."
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}

	for _,item := range items{
		if item.Name == input.name{
			response.status = "Item name already exist..."
			ctx.IndentedJSON(http.StatusNotFound,response)
			return
		}
	}
	item:=entity.Item{
		Name:input.name,
		Brand:input.brand,
		Description: input.description,
		Image: input.imageurl,
		Price: input.price,
		Number: input.number,
		ProductionDate:input.production_date,
		ExpireDate : input.expire_date,
	}


	itm,ers:= itemHandler.itemServ.CreateItem(item)
	if len(ers) >0{
		response.status = "Internal Server Error ..."
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	response.status = "Item create successful"
	response.item = itm

	ctx.IndentedJSON(http.StatusCreated,response)

}

func (itemHandler *ItemHandler)UpdateItem(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		item *entity.Item
	}{
		status:"Unauthorized user",
	}
	input := &struct{
		name string
		description string
		brand string
		imageurl string
		price float64
		production_date time.Time
		expire_date time.Time
	}{}
	sess := itemHandler.sessionHa.GetSession(ctx)
	if sess != nil || sess.Role !="Admin"{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	e := ctx.BindJSON(&input)
	if e != nil{
		ctx.JSON(404,e)
	}

	if input.name == "" || input.description== "" || input.brand == "" || input.imageurl == "" || input.production_date.String() =="" || input.expire_date.String() == "" || input.price <= 0{
		response.status = "Incorrect input..."
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}
	id,_:= strconv.Atoi(ctx.Param("id"))
	
	item,ers := itemHandler.itemServ.GetItem(uint(id))
	
	if item.Name != input.name{
		if itemHandler.itemServ.IsItemNameExist(input.name) != nil{
			response.status = "Item name already exist..."
			ctx.IndentedJSON(http.StatusNotFound,response)
			return
		}
		
	}

	itm:=entity.Item{
		Name:input.name,
		Brand:input.brand,
		Description: input.description,
		Image: input.imageurl,
		Price: input.price,
		ProductionDate:input.production_date,
		ExpireDate : input.expire_date,
	}
	item.ID = uint(id)
	item,ers = itemHandler.itemServ.UpdateItem(itm)
	if ers != nil{
		response.status = "Internal Server Error ..."
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	response.status = "upadate succeful"
	response.item = item
	ctx.JSON(200,response)
}

func (itemHandler *ItemHandler)DeleteItem(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		item *entity.Item
	}{
		status:"Unauthorized user",
	}
	sess := itemHandler.sessionHa.GetSession(ctx)
	if sess != nil || sess.Role !="Admin"{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	
	id,e:= strconv.Atoi(ctx.Param("id"))
	if e != nil{
		response.status = "Incorrect Format..." 
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	itm,ers:= itemHandler.itemServ.GetItem(uint(id))
	if len(ers)>0{
		response.status = "No such user..." 
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	
	itm,ers = itemHandler.itemServ.DeleteItem(uint(id))
	if ers != nil{
		response.status = "Internal Server Error ..."
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	catagories,ers:= itemHandler.cataSrv.GetCatagories()
	if ers != nil{
		response.status = "Internal Server Error ..."
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}

	for _,cata := range catagories{
		check := false
		for _,itm_Id := range cata.Items_Id{
			if itm_Id != itm.ID{
				cata.Items_Id = append(cata.Items_Id, itm_Id)
			}else{
				check = true
			}
		}
		if check{
			_,ers := itemHandler.cataSrv.UpdateCatagory(cata)
			if len(ers)>0{
				response.status = "Internal Server Error ..."
				ctx.IndentedJSON(http.StatusInternalServerError,response)
				return
			}
		}
	}
	response.status = "Delete successful"
	response.item = itm
	ctx.JSON(200,response)
}
