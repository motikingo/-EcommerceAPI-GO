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

	recordRepository "github.com/motikingo/ecommerceRESTAPI-Go/record/repository"
	recordService "github.com/motikingo/ecommerceRESTAPI-Go/record/service"

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
	defer db.Close()
	r := gin.Default()
	session := handler.NewSessionHandler()
	secure := middleware.NewMiddlerware(&session) 
	admRepo := repository.NewUserRepo(db)
	admServc := service.NewUserSrvc(admRepo)
	admHandler:= handler.NewAdminHandler(admServc,&session)
	
	r.GET("/admins",secure.AdminLogedIn(),admHandler.GetAdmins)
	r.GET("/admin/:id",secure.AdminLogedIn(),admHandler.GetAdmin)
	r.POST("/create/admin",admHandler.CreateAdmin)
	r.POST("/admin/Login",admHandler.AdminLogIn)
	r.GET("/admin/Logout",secure.AdminLogedIn(),admHandler.AdminLogOut)

	r.PUT("/ChangeProfile/admin",secure.AdminLogedIn(),admHandler.ChangeProfile)
	r.PUT("/ChangePassword/admin",secure.AdminLogedIn(),admHandler.ChangePassword)

	userRepo := repository.NewUserRepo(db)
	userServc := service.NewUserSrvc(userRepo)

	itemRepo := itemRepository.NewItemRepo(db)
	itemServ:= itemService.NewItemServ(itemRepo)

	catRepo := catagoryRepository.NewCatagoryRepo(db)
	catservs:= catagoryService.NewCatagoryServ(catRepo)

	recordrepo := recordRepository.NewRecordRepo(db)
	recordSrv := recordService.NewRecordServ(recordrepo)

	cartHa := handler.NewcartHandler(recordSrv,&session)
	userHandler:= handler.NewUserHandler(userServc,itemServ,&cartHa,recordSrv,&session)

	recoHa := handler.NewRecordHandler(recordSrv,cartHa,&session)

	r.GET("/users",secure.AdminLogedIn(),userHandler.GetUsers)
	r.GET("/user/:id",secure.UserLogedIn(),userHandler.GetUser)

	r.POST("/create/user",userHandler.CreateUser)

	r.POST("/user/Login",userHandler.LogIn)
	r.GET("/user/Logout",secure.UserLogedIn(),userHandler.Logout)

	r.PUT("/ChangeProfile/",secure.UserLogedIn(),userHandler.ChangeProfile)
	r.PUT("/ChangePassword/",secure.UserLogedIn(),userHandler.ChangePassword)
	r.GET("/CreateCart/user/",secure.UserLogedIn(),cartHa.CreateCart)
	r.POST("/AddItem/user/",secure.UserLogedIn(),userHandler.AddItemToMyCart)
	r.PUT("/DeleteItemFromCart/user/:item_id",secure.UserLogedIn(),userHandler.DeleteItemFromMyCart)
	r.GET("/DeleteMyCart/user/",secure.UserLogedIn(),userHandler.DeleteMyCart)
	r.GET("/updateItemInMyCart/user/:item_id",secure.UserLogedIn(),userHandler.AddItemToMyCart)

	r.GET("/Order/user",secure.UserLogedIn(),userHandler.Order)
	r.GET("/MyRecord/user",secure.UserLogedIn(),recoHa.GetRecord)
	r.GET("/ClearRecord/user",secure.UserLogedIn(),recoHa.ClearRecord)


	
	r.DELETE("/delete/",secure.UserLogedIn(),userHandler.DeleteAccount)

	catHandler := handler.NewcatHandler(catservs,itemServ,&session)

	r.GET("/catagories",catHandler.GetCatagories)
	r.GET("/catagory/:id",catHandler.GetCatagory)
	r.PUT("/catagory/update/:id",secure.AdminLogedIn(), catHandler.UpdateCatagory)
	r.POST("/catagory/create/",secure.AdminLogedIn(),catHandler.CreateCatagory)
	r.DELETE("/catagory/delete/:id",secure.AdminLogedIn(),catHandler.DeleteCatagory)
	r.GET("/catagory/Items",secure.AdminLogedIn(),catHandler.GetMyItems)
	r.POST("/catagory/AddItems/:id",secure.AdminLogedIn(),catHandler.AddItems)
	r.PUT("catagory/DeleteItem/:id",secure.AdminLogedIn(),catHandler.DeleteItemFromCatagory)

	itemHandler:= handler.NewItemHandler(itemServ,catservs,&session)

	r.GET("/items",itemHandler.GetItems)
	r.GET("/item/:id",itemHandler.GetItem)
	r.PUT("/update/item/:id",secure.AdminLogedIn(),itemHandler.UpdateItem)
	r.POST("/create/item/",secure.AdminLogedIn(),itemHandler.CreateItem)
	r.DELETE("/delete/item/:id",secure.AdminLogedIn(),itemHandler.DeleteItem)

	recordHandler := handler.NewRecordHandler(recordSrv,cartHa,&session)

	r.GET("/records",secure.AdminLogedIn(),recordHandler.GetRecords)
	r.GET("/Myrecord/",secure.UserLogedIn(),recordHandler.GetRecord)
	//r.POST("/create/record/",secure.UserLogedIn(),recordHandler.CreateCart)
	r.DELETE("/delete/record/",secure.UserLogedIn(),recordHandler.ClearRecord)
	
	log.Fatal(r.Run(":80"))

}

