package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"

	cartRepository "github.com/motikingo/ecommerceRESTAPI-Go/cart/repository"
	cartService "github.com/motikingo/ecommerceRESTAPI-Go/cart/service"
	catagoryRepository "github.com/motikingo/ecommerceRESTAPI-Go/catagory/repository"
	catagoryService "github.com/motikingo/ecommerceRESTAPI-Go/catagory/service"
	"github.com/motikingo/ecommerceRESTAPI-Go/user/repository"
	"github.com/motikingo/ecommerceRESTAPI-Go/user/service"

	itemRepository "github.com/motikingo/ecommerceRESTAPI-Go/item/repository"
	itemService "github.com/motikingo/ecommerceRESTAPI-Go/item/service"

	"github.com/motikingo/ecommerceRESTAPI-Go/handler"

	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

func main(){
	err = godotenv.Load("../.env")

	if err !=nil{
		log.Fatal(err)
		return
	}

	dialect := os.Getenv("dialect")
	dbname  := os.Getenv("dbname")
	host := os.Getenv("host")
	user := os.Getenv("user")
	password:= os.Getenv("password")

	dbURL := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode = disable",host,user,dbname,password)

	db,err = gorm.Open(dialect,dbURL) 

	if err !=nil{
		log.Fatal(err)
		return
	}

	db.Debug().AutoMigrate(entity.User{})
	db.Debug().AutoMigrate(entity.Catagory{})
	db.Debug().AutoMigrate(entity.Item{})
	db.Debug().AutoMigrate(entity.Cart{})

	r := gin.Default()

	userRepo:=repository.NewUserRepo(db)
	userServc := service.NewUserSrvc(userRepo)
	userHandler:= handler.NewUserHandler(userServc)

	r.GET("/users",userHandler.GetUsers)
	r.GET("/users/:id",userHandler.GetUser)
	r.POST("/create/user/",userHandler.CreateUser)
	r.PUT("/update/user/:id",userHandler.UpdateUser)
	r.POST("/delete/user/",userHandler.DeleteUser)

	cartRepo:= cartRepository.NewCartRepo(db)
	cartServ := cartService.NewCartServ(cartRepo)
	cartHadler := handler.NewcartHandler(cartServ) 

	r.GET("/{user_id}/carts/",cartHadler.GetCarts)
	r.GET("/{user_id}/carts/:id",cartHadler.GetCart)
	r.PUT("/{user_id}/update/carts/:id",cartHadler.UpdateCart)
	r.POST("/{user_id}/create/carts/",cartHadler.CreateCarts)
	r.DELETE("/{use_id}/delete/carts/:id",cartHadler.DeleteCarts)

	catRepo := catagoryRepository.NewCatagoryRepo(db)
	catservs:= catagoryService.NewCatagoryServ(catRepo)
	catHandler := handler.NewcatHandler(catservs)

	r.GET("/catagories",catHandler.GetCatagories)
	r.GET("/catagory/:id",catHandler.GetCatagory)
	r.PUT("/update/catagory/:id",catHandler.UpdateCatagory)
	r.POST("/create/catagory/",catHandler.CreateCatagory)
	r.DELETE("/delete/catagory/:id",catHandler.DeleteCatagory)


	itemRepo := itemRepository.NewItemRepo(db)
	itemServ:= itemService.NewItemServ(itemRepo)
	itemHandler:= handler.NewItemHandler(itemServ)

	r.GET("/items",itemHandler.GetItems)
	r.GET("/item/:id",itemHandler.GetItem)
	r.PUT("/update/item/:id",itemHandler.UpdateItem)
	r.POST("/create/item/",itemHandler.CreateItem)
	r.DELETE("/delete/item/:id",itemHandler.DeleteItem)
	
	log.Fatal(r.Run(":80"))

}








