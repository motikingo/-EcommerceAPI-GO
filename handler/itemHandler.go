package handler

import (
	//"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	"github.com/motikingo/ecommerceRESTAPI-Go/item"
	"github.com/motikingo/ecommerceRESTAPI-Go/catagory"

)

type ItemHandler struct{
	itemServ item.ItemService
	sessioHa *SessionHandler
	cataSrv catagory.CatagoryService
}

func NewItemHandler(itemServ item.ItemService)ItemHandler{
	return ItemHandler{itemServ: itemServ}
}

func (itemHandler *ItemHandler)GetItems(ctx *gin.Context){
	cxt.Header("Content-Type","application/json")

	response := &struct{
		status string
		items [] entity.Item
	}{
		status:"Unauthorized user"
	}

	sess := ursHandler.sessionHa.GetSession(r)
	if sess != nil {
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}

	items,ers:=itemHandler.itemServ.GetItems()
	if  len (ers)>0 {
		response.status = "Internal Server Error"
		cxt.IndentedJSON(http.StatusInternalServerError,response)
		return
	}

	if len(items)==0 {
		response.status = "No item added yet"
		cxt.IndentedJSON(http.StatusOK,response)
		return
	}
	response.status = "items retrieved successfully"
	response.items = items
	ctx.JSON(200,items)


}
func (itemHandler *ItemHandler)GetItem(ctx *gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		items [] entity.Item
	}{
		status:"Unauthorized user"
	}

	sess := ursHandler.sessionHa.GetSession(r)
	if sess != nil {
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}


	id,e:= strconv.Atoi(ctx.Param("id")) 
	if  e != nil {
		response.status = "bad request..."
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	items,ers:=itemHandler.itemServ.GetItem(uint(id))
	if  len (ers)>0 {
		response.status = "No such Item"
		cxt.IndentedJSON(http.StatusNotFound,response)
		return
	}
	response.status = "Item fetched"
	cxt.IndentedJSON(http.StatusOK,response)
	ctx.IndentedJSON(200,response)

}

func (itemHandler *ItemHandler)CreateItem(ctx *gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		item *entity.Item
	}{
		status:"Unauthorized user"
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
	sess := ursHandler.sessionHa.GetSession(r)
	if sess != nil || sess.Role !="Admin"{
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}

	e := ctx.BindJSON(&input)
	if e != nil{
		response.status = "failed to bind input"
		cxt.IndentedJSON(http.StatusNotFound,response)
		return
	}
	if input.name == "" || input.description== "" || input.brand == "" || input.imageurl == "" || input.production_date ==nil || input.expire_date== nil || input.price <= 0 || number <=0{
		response.status = "Incorrect input..."
		cxt.IndentedJSON(http.StatusNotFound,response)
		return
	}

	if input.expire_date < time.Now(){
		response.status = "this Item is expired..."
		cxt.IndentedJSON(http.StatusNotFound,response)
		return
	}
	items,ers := itemHandler.itemServ.GetItems()

	if len(ers)>0{
		response.status = "Internal Server Error ..."
		cxt.IndentedJSON(http.StatusInternalServerError,response)
		return
	}

	for _,item := range items{
		if item.Name == input.name{
			response.status = "Item name already exist..."
			cxt.IndentedJSON(http.StatusNotFound,response)
			return
		}
	}
	item:=entity.Item{
		Name:input.name
		Brand:input.brand
		Description: input.description
		Image: input.imageurl
		Price: input.price
		Number: input.number
		ProductionDate:input.production_date
		ExpireDate : input.expire_date
	}


	itm,ers:= itemHandler.itemServ.CreateItem(item)
	if len(ers) >0{
		response.status = "Internal Server Error ..."
		cxt.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	response.status = "Item create successful"
	response.item = itm

	ctx.IndentedJSON(http.StatusCreated,response)

}

func (itemHandler *ItemHandler)UpdateItem(ctx *gin.Context){
	cxt.Header("Content-Type","application/json")
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		item *entity.Item
		catagorySrv *
	}{
		status:"Unauthorized user"
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
	sess := ursHandler.sessionHa.GetSession(r)
	if sess != nil || sess.Role !="Admin"{
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}
	id,_:= strconv.Atoi(ctx.Param("id"))
	
	e = ctx.BindJSON(&input)
	if e != nil{
		ctx.JSON(404,e)
	}

	if input.name == "" || input.description== "" || input.brand == "" || input.imageurl == "" || input.production_date ==nil || input.expire_date== nil || input.price <= 0 || number <=0{
		response.status = "Incorrect input..."
		cxt.IndentedJSON(http.StatusNotFound,response)
		return
	}

	if len(ers)>0{
		response.status = "Internal Server Error ..."
		cxt.IndentedJSON(http.StatusInternalServerError,response)
		return
	}

	for _,item := range items{
		if item.Name == input.name{
			response.status = "Item name already exist..."
			cxt.IndentedJSON(http.StatusNotFound,response)
			return
		}
	}

	item:=entity.Item{
		Name:input.name
		Brand:input.brand
		Description: input.description
		Image: input.imageurl
		Price: input.price
		ProductionDate:input.production_date
		ExpireDate : input.expire_date
	}
	item.ID = uint(id)
	itm,ers:= itemHandler.itemServ.UpdateItem(item)
	if ers != nil{
		response.status = "Internal Server Error ..."
		cxt.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	response.status = "upadate succeful"
	response.item = itm
	ctx.JSON(200,response)
}

func (itemHandler *ItemHandler)DeleteItem(ctx *gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		item *entity.Item
	}{
		status:"Unauthorized user"
	}

	if sess != nil || sess.Role !="Admin"{
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}
	var item_Id string
	id,e:= strconv.Atoi(ctx.BindJSON(item_Id))
	if e != nil || id == nil{
		response.status = "Incorrect Format..." 
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	itm,ers:= itemHandler.itemServ.Item(uint(id))
	if len(ers)>0{
		response.status = "No such user..." 
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}
	
	itm,ers:= itemHandler.itemServ.DeleteItem(uint(id))
	if ers != nil{
		response.status = "Internal Server Error ..."
		cxt.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	catagories,ers:= itemHandler.cataSrv.GetCatagories()
	if ers != nil{
		response.status = "Internal Server Error ..."
		cxt.IndentedJSON(http.StatusInternalServerError,response)
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
			cat,ers := itemHandler.cataSrv.UpdateCatagory(cata)
			if len(ers)>0{
				response.status = "Internal Server Error ..."
				cxt.IndentedJSON(http.StatusInternalServerError,response)
				return
			}
		}
	}
	response.status = "Delete successful"
	response.item = itm
	ctx.JSON(200,response)
}
