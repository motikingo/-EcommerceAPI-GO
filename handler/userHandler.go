package handler

import (
	//"fmt"
	"log"
	//"net/http"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	"github.com/motikingo/ecommerceRESTAPI-Go/user"
)

type UserHandler struct {

	userSrv user.UserService
}

func NewUserHandler(userSrv user.UserService)UserHandler{
	return UserHandler{userSrv: userSrv}
}

func (usrHandler *UserHandler)GetUsers(cxt *gin.Context){

	users,e:= usrHandler.userSrv.GetUsers()

	if len(e)>0{
		log.Fatal(e)
	}

	usr,er:= json.MarshalIndent(users,"","/t/t")
	if er!=nil{
		log.Fatal(e)
	}
	cxt.JSON(200,string(usr))

}

func (usrHandler *UserHandler)GetUser(cxt *gin.Context){
	cxt.Header("Content-Type","application/json")
	id,err:= strconv.Atoi(cxt.Param("id"))
	if err!= nil{
		log.Fatal(err)
	}
	user,e:= usrHandler.userSrv.GetUser(uint(id))

	if len(e)>0{
		log.Fatal(e)
	}

	cxt.JSON(200,user)

}

func(usrHandler *UserHandler) CreateUser(cxt *gin.Context){
	cxt.Header("Content-Type","application/json")
	var	usr entity.User

	cxt.BindJSON(&usr)
	usrHandler.userSrv.CreateUser(entity.User(usr))

	user,err:=  json.MarshalIndent(usr,"","/r/r")

	if err!=nil {
		log.Fatal(err)
	}

	cxt.JSON(200,string(user))
}

func(usrHandler *UserHandler) UpdateUser(cxt * gin.Context){
	cxt.Header("Content-Type","application/json")
	var usr entity.User
	id,e:=strconv.Atoi(cxt.Param("id"))
	if e!=nil {
		log.Fatal(e)
	}
	err := cxt.BindJSON(&usr)

	if err!=nil {
		
		log.Fatal("it,s here")
	}

	user,ers:=usrHandler.userSrv.UpdateUser(uint(id),usr)

	if ers!=nil {
		log.Fatal(ers)
	}
	ursMarsh,_:=json.MarshalIndent(user,"","/t/")

	cxt.JSON(200,ursMarsh)




	

}

func(usrHandler *UserHandler) DeleteUser(cxt * gin.Context){

}

