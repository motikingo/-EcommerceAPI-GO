package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"

	"github.com/motikingo/ecommerceRESTAPI-Go/user/repository"
	"github.com/motikingo/ecommerceRESTAPI-Go/user/service"

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
	
	userRepo:=repository.NewUserRepo(db)
	userServc := service.NewUserSrvc(userRepo)
	userHandler:= handler.NewUserHandler(userServc)
	
	r := gin.Default()

	r.GET("/users",userHandler.GetUsers)

	r.GET("/users/:id",userHandler.GetUser)
	r.POST("/create/user/",userHandler.CreateUser)
	r.PUT("/update/user/:id",userHandler.UpdateUser)
	r.POST("/delete/user/",userHandler.DeleteUser)

	r.Run(":80")

}

// func FirstTry(cxt *gin.Context){
// 	cxt.JSON(200,gin.H{
// 		"message":"OK",
// 	})
// }







