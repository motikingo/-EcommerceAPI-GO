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

type ItemHandler struct {
	itemServ  item.ItemService
	cataSrv   catagory.CatagoryService
	sessionHa *SessionHandler
}

func NewItemHandler(itemServ item.ItemService, cataSrv catagory.CatagoryService, sessionHa *SessionHandler) ItemHandler {
	return ItemHandler{itemServ: itemServ, cataSrv: cataSrv, sessionHa: sessionHa}
}

func (itemHandler *ItemHandler) GetItems(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	response := &struct {
		Status string
		Items  []entity.Item
	}{
		Status: "Unauthorized user",
	}

	sess := itemHandler.sessionHa.GetSession(ctx)
	if sess == nil {
		ctx.IndentedJSON(http.StatusUnauthorized, response)
		return
	}

	items, ers := itemHandler.itemServ.GetItems()
	if len(ers) > 0 {
		response.Status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusInternalServerError, response)
		return
	}

	if len(items) == 0 {
		response.Status = "No item added yet"
		ctx.IndentedJSON(http.StatusOK, response)
		return
	}
	response.Status = "items retrieved successfully"
	response.Items = items
	ctx.JSON(200, items)

}
func (itemHandler *ItemHandler) GetItem(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	response := &struct {
		Status string
		Item   *entity.Item
	}{
		Status: "Unauthorized user",
	}

	sess := itemHandler.sessionHa.GetSession(ctx)
	if sess == nil {
		ctx.IndentedJSON(http.StatusUnauthorized, response)
		return
	}

	id, e := strconv.Atoi(ctx.Param("id"))
	if e != nil {
		response.Status = "bad request..."
		ctx.IndentedJSON(http.StatusBadRequest, response)
		return
	}

	item, ers := itemHandler.itemServ.GetItem(uint(id))
	if len(ers) > 0 {
		response.Status = "No such Item"
		ctx.IndentedJSON(http.StatusNotFound, response)
		return
	}
	response.Status = "Item fetched"
	response.Item = item
	ctx.IndentedJSON(http.StatusOK, response)
}

func (itemHandler *ItemHandler) CreateItem(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	response := &struct {
		Status string
		Item   entity.Item
	}{
		Status: "Unauthorized user",
	}
	input := &struct {
		Name        string
		Description string
		Brand       string
		Imageurl    string
		Price       float64
		Number      int
		// Production_date time.Time
		// Expire_date time.Time
	}{}
	sess := itemHandler.sessionHa.GetSession(ctx)
	if sess == nil || sess.Role != "Admin" {
		ctx.IndentedJSON(http.StatusUnauthorized, response)
		return
	}

	if e := ctx.BindJSON(&input); e != nil {
		response.Status = "failed to bind input"
		ctx.IndentedJSON(http.StatusNotFound, response)
		return
	}
	if input.Name == "" || input.Description == "" || input.Brand == "" || input.Imageurl == "" || input.Price <= 0 || input.Number <= 0 {
		response.Status = "Incorrect input..."
		ctx.IndentedJSON(http.StatusNotFound, response)
		return
	}

	// if input.Expire_date.Before(time.Now()){
	// 	response.Status = "this Item is expired..."
	// 	ctx.IndentedJSON(http.StatusNotFound,response)
	// 	return
	// }
	// items,ers := itemHandler.itemServ.GetItems()

	// if len(ers)>0{
	// 	response.Status = "Internal Server Error ..."
	// 	ctx.IndentedJSON(http.StatusInternalServerError,response)
	// 	return
	// }

	// for _,item := range items{
	// 	if item.Name == input.Name{
	// 		response.Status = "Item name already exist..."
	// 		ctx.IndentedJSON(http.StatusNotFound,response)
	// 		return
	// 	}
	// }
	if itemHandler.itemServ.IsItemNameExist(input.Name) != nil {
		response.Status = "Item name already exist..."
		ctx.IndentedJSON(http.StatusNotFound, response)
		return
	}
	item := entity.Item{
		Name:           input.Name,
		Description:    input.Description,
		Brand:          input.Brand,
		Image:          input.Imageurl,
		Price:          input.Price,
		Number:         input.Number,
		ProductionDate: time.Now().AddDate(-1, 5, 12),
		ExpireDate:     time.Now().AddDate(1, 2, 5),
	}
	itm, ers := itemHandler.itemServ.CreateItem(item)
	if len(ers) > 0 {
		response.Status = "Internal Server Error ..."
		ctx.IndentedJSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = "Item create successful"
	response.Item = *itm

	ctx.IndentedJSON(http.StatusCreated, response)

}

func (itemHandler *ItemHandler) UpdateItem(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	response := &struct {
		Status string
		Item   *entity.Item
	}{
		Status: "Unauthorized user",
	}
	input := &struct {
		Name        string
		Description string
		Brand       string
		Imageurl    string
		Price       float64
	}{}
	sess := itemHandler.sessionHa.GetSession(ctx)
	if sess == nil || sess.Role != "Admin" {
		ctx.IndentedJSON(http.StatusUnauthorized, response)
		return
	}
	e := ctx.BindJSON(&input)
	if e != nil {
		ctx.JSON(404, e)
	}

	if input.Name == "" || input.Description == "" || input.Brand == "" || input.Imageurl == "" || input.Price <= 0 {
		response.Status = "Incorrect input..."
		ctx.IndentedJSON(http.StatusNotFound, response)
		return
	}
	id, _ := strconv.Atoi(ctx.Param("id"))

	item, ers := itemHandler.itemServ.GetItem(uint(id))

	if len(ers) > 0 {
		response.Status = "Internal Server Error ..."
		ctx.IndentedJSON(http.StatusInternalServerError, response)
		return
	}

	if item.Name != input.Name {
		if itemHandler.itemServ.IsItemNameExist(input.Name) != nil {
			response.Status = "Item name already exist..."
			ctx.IndentedJSON(http.StatusNotFound, response)
			return
		}

	}

	itm := entity.Item{
		Name:        input.Name,
		Brand:       input.Brand,
		Description: input.Description,
		Image:       input.Imageurl,
		Price:       input.Price,
	}
	itm.ID = uint(id)
	item, ers = itemHandler.itemServ.UpdateItem(itm)
	if ers != nil {
		response.Status = "Internal Server Error ..."
		ctx.IndentedJSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = "upadate succeful"
	response.Item = item
	ctx.JSON(200, response)
}

func (itemHandler *ItemHandler) DeleteItem(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	response := &struct {
		Status string
		Item   *entity.Item
	}{
		Status: "Unauthorized user",
	}
	sess := itemHandler.sessionHa.GetSession(ctx)
	if sess == nil || sess.Role != "Admin" {
		ctx.IndentedJSON(http.StatusUnauthorized, response)
		return
	}

	id, e := strconv.Atoi(ctx.Param("id"))
	if e != nil {
		response.Status = "Incorrect Format..."
		ctx.IndentedJSON(http.StatusBadRequest, response)
		return
	}

	item, ers := itemHandler.itemServ.GetItem(uint(id))

	if len(ers) > 0 || item == nil {
		response.Status = "No such Item..."
		ctx.IndentedJSON(http.StatusBadRequest, response)
		return
	}

	item, ers = itemHandler.itemServ.DeleteItem(uint(id))
	if len(ers) > 0 {
		response.Status = "Internal Server Error ..."
		ctx.IndentedJSON(http.StatusInternalServerError, response)
		return
	}
	// catagories,ers:= itemHandler.cataSrv.GetCatagories()
	// if len(ers) >0{
	// 	response.Status = "No catagory Added yet..."
	// 	ctx.IndentedJSON(http.StatusInternalServerError,response)
	// 	return
	// }

	// for _,cata := range catagories{
	// 	check := false
	// 	for _,itm := range cata.Items{
	// 		if item.ID != itm.ID{
	// 			cata.Items = append(cata.Items, itm)
	// 		}else{
	// 			check = true
	// 		}
	// 	}
	// 	if check{
	// 		_,ers := itemHandler.cataSrv.UpdateCatagory(cata)
	// 		if len(ers)>0{
	// 			response.Status = "Internal Server Error ..."
	// 			ctx.IndentedJSON(http.StatusInternalServerError,response)
	// 			return
	// 		}
	// 	}
	// }
	response.Status = "Delete successful"
	response.Item = item
	ctx.JSON(200, response)
}
