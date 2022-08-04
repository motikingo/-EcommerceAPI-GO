package handler

import (
	//"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	"github.com/motikingo/ecommerceRESTAPI-Go/item"
)

type ItemHandler struct{
	itemServ item.ItemService
}

func NewItemHandler(itemServ item.ItemService)ItemHandler{
	return ItemHandler{itemServ: itemServ}
}

func (itemHandler *ItemHandler)GetItems(ctx *gin.Context){
	items,ers:=itemHandler.itemServ.GetItems()
	if  len (ers)>0 {
		ctx.JSON(404,ers)
	}

	ctx.JSON(200,items)


}
func (itemHandler *ItemHandler)GetItem(ctx *gin.Context){

	id,e:= strconv.Atoi(ctx.Param("id")) 
	if  e != nil {
		ctx.JSON(404,e)
	}

	items,ers:=itemHandler.itemServ.GetItem(uint(id))
	if  len (ers)>0 {
		ctx.JSON(404,ers)
	}

	ctx.JSON(200,items)

}
func (itemHandler *ItemHandler)UpdateItem(ctx *gin.Context){
	var item entity.Item
	id,e:= strconv.Atoi(ctx.Param("id"))
	if e != nil{
		ctx.JSON(404,e)
	}
	e = ctx.BindJSON(&item)
	if e != nil{
		ctx.JSON(404,e)
	}
	itm,ers:= itemHandler.itemServ.UpdateItem(uint(id),item)
	if ers != nil{
		ctx.JSON(404,ers)
	}
	ctx.JSON(200,itm)
}
func (itemHandler *ItemHandler)CreateItem(ctx *gin.Context){
	var item entity.Item
	e := ctx.BindJSON(&item)
	if e != nil{
		ctx.JSON(404,e)
	}
	itm,ers:= itemHandler.itemServ.CreateItem(item)
	if ers != nil{
		ctx.JSON(404,ers)
	}
	ctx.JSON(200,itm)

}
func (itemHandler *ItemHandler)DeleteItem(ctx *gin.Context){
	id,e:= strconv.Atoi(ctx.Param("id"))
	if e != nil{
		ctx.JSON(404,e)
	}
	itm,ers:= itemHandler.itemServ.DeleteItem(uint(id))
	if ers != nil{
		ctx.JSON(404,ers)
	}
	ctx.JSON(200,itm)
}
