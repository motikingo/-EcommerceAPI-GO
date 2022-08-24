package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	catagoryRepository "github.com/motikingo/ecommerceRESTAPI-Go/catagory/repository"
	catagoryService "github.com/motikingo/ecommerceRESTAPI-Go/catagory/service"
	"github.com/motikingo/ecommerceRESTAPI-Go/customer/repository"
	"github.com/motikingo/ecommerceRESTAPI-Go/customer/service"

	itemRepository "github.com/motikingo/ecommerceRESTAPI-Go/item/repository"
	itemService "github.com/motikingo/ecommerceRESTAPI-Go/item/service"

	cartRepository "github.com/motikingo/ecommerceRESTAPI-Go/cart/repository"
	cartService "github.com/motikingo/ecommerceRESTAPI-Go/cart/service"

	"github.com/motikingo/ecommerceRESTAPI-Go/database"
	"github.com/motikingo/ecommerceRESTAPI-Go/handler"
	"github.com/motikingo/ecommerceRESTAPI-Go/middleware"
)

var db *gorm.DB

func init(){

	db = database.Connect()
	if db == nil {
		return
	}
	database.MigrateModel(db)
}

func main(){	
	r := gin.Default()

	session := handler.NewSessionHandler()

	secure := middleware.NewMiddlerware(&session) 

	admRepo := repository.NewUserRepo(db)
	admServc := service.NewUserSrvc(admRepo)
	admHandler:= handler.NewAdminHandler(admServc,&session)
	
	r.GET("/admins",secure.AdminLogedIn(),admHandler.GetAdmins)
	r.GET("/admin/:id",secure.AdminLogedIn(),admHandler.GetAdmin)
	r.POST("/create/admin",secure.AdminLogedIn(),admHandler.CreateAdmin)
	r.PUT("/ChangeProfile/admin",secure.AdminLogedIn(),admHandler.ChangeProfile)
	r.PUT("/ChangePassword/admin",secure.AdminLogedIn(),admHandler.ChangePassword)

	userRepo := repository.NewUserRepo(db)
	userServc := service.NewUserSrvc(userRepo)
	userHandler:= handler.NewUserHandler(userServc,&session)
	
	r.GET("/users",secure.AdminLogedIn(),userHandler.GetUsers)
	r.GET("/user:id",secure.UserLogedIn(),userHandler.GetUser)
	r.POST("/create/",secure.UserLogedIn(),userHandler.CreateUser)
	r.PUT("/ChangeProfile/",secure.UserLogedIn(),userHandler.ChangeProfile)
	r.PUT("/ChangePassword/",secure.UserLogedIn(),userHandler.ChangePassword)
	r.GET("/AddItem/user/",secure.UserLogedIn(),userHandler.AddItemToMyCart)
	r.GET("/Order/",secure.UserLogedIn(),userHandler.Order)
	r.DELETE("/delete/",secure.UserLogedIn(),userHandler.DeleteAccount)

	catRepo := catagoryRepository.NewCatagoryRepo(db)
	catservs:= catagoryService.NewCatagoryServ(catRepo)
	catHandler := handler.NewcatHandler(catservs,&session)

	r.GET("/catagories",catHandler.GetCatagories)
	r.GET("/catagory/:id",catHandler.GetCatagory)
	r.PUT("catagory/update/",secure.AdminLogedIn(), catHandler.UpdateCatagory)
	r.POST("catagory/create/",secure.AdminLogedIn(),catHandler.CreateCatagory)
	r.DELETE("catagory/delete/",secure.AdminLogedIn(),catHandler.DeleteCatagory)
	r.GET("/catagory/Items",secure.AdminLogedIn(),catHandler.GetMyItems)
	r.POST("/catagory/AddItems/",secure.AdminLogedIn(),catHandler.AddItems)
	r.PUT("catagory/DeleteItem/",secure.AdminLogedIn(),catHandler.DeleteItemFromCatagory)

	itemRepo := itemRepository.NewItemRepo(db)
	itemServ:= itemService.NewItemServ(itemRepo)
	itemHandler:= handler.NewItemHandler(itemServ,&session)

	r.GET("/items",itemHandler.GetItems)
	r.GET("/item/:id",itemHandler.GetItem)
	r.PUT("/update/item/",secure.AdminLogedIn(),itemHandler.UpdateItem)
	r.POST("/create/item/",secure.AdminLogedIn(),itemHandler.CreateItem)
	r.DELETE("/delete/item/",secure.AdminLogedIn(),itemHandler.DeleteItem)

	cartrepo := cartRepository.NewCartRepo(db)
	cartSrv := cartService.NewCartServ(cartrepo)
	cartHandler := handler.NewcartHandler(cartSrv,&session)

	r.GET("/carts",secure.AdminLogedIn(),cartHandler.GetCarts)
	r.GET("/cart/:id",secure.UserLogedIn(),cartHandler.GetCart)
	r.POST("/create/cart/",secure.UserLogedIn(),cartHandler.CreateCart)
	r.DELETE("/delete/cart/",secure.UserLogedIn(),cartHandler.DeleteCart)
	
	log.Fatal(r.Run(":80"))

}

