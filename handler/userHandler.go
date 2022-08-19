package handler

import (
	//"fmt"
	"log"
	//"net/http"
	//"encoding/json"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	"github.com/motikingo/ecommerceRESTAPI-Go/customer"
)

type UserHandler struct {

	userSrv user.UserService
	sessionHa *SessionHandler
}

func NewUserHandler(userSrv user.UserService, sessionHa *SessionHandler)UserHandler{
	return UserHandler{userSrv: userSrv,sessionHa:sessionHa}
}

func (usrHandler *UserHandler)GetUsers(cxt *gin.Context){
	cxt.Header("Content-Type","application/json")
	sess := ursHandler.sessionHa.GetSession(r)
	if sess != nil || sess.Role != "Admin"{
		cxt.IndentedJSON(http.StatusUnathorized,"Unauthorized user")
		return
	}
	users,e:= usrHandler.userSrv.GetUsers()

	if len(e)>0 {
		cxt.IndentedJSON(http.InternalServerError,"Internal Server Error")
		return
	}
	if users == nil{
		cxt.IndentedJSON(http.StatusOk,"No user Found")
		return
	}

	cxt.IndentedJSON(200,users)

}

func (usrHandler *UserHandler)GetUser(cxt *gin.Context){
	cxt.Header("Content-Type","application/json")
	sess := ursHandler.sessionHa.GetSession(r)
	if sess != nil || sess.Role != "Admin"{
		cxt.IndentedJSON(http.StatusUnathorized,"Unauthorized user")
		return
	}

	user,e:= usrHandler.userSrv.GetUser()

	if len(e)>0{
		cxt.IndentedJSON(http.InternalServerError,"Internal Server Error")
		return

	}

	cxt.JSON(200,user)

}

func(usrHandler *UserHandler) CreateUser(cxt *gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		user *entity.User
	}{
		status:"Unauthorized user"
	}
	input := &struct{
		name string
		last_name string		
		user_name string
		email string
		password string
		comfirm_password string


	}{}

	if er = cxt.BindJSON(&input); er!=nil{
		response.status = "bad request..." 
		cxt.IndentedJSON(cxt.StatusBadRequest,response)
		return
	}

	input.password = helper.SecurePassword(input.password)

	if input.UserName =="" || input.Email == "" || input.Name=="" || input.LastName || input.password == "" || input.comfirm_password == ""{
		response.status = "Invalid input..."
		cxt.IndentedJSON(cxt.StatusBadRequest,response)
		return
	}

	users,er := userHandler.userSrv.GetUsers()

	if len(er)>0 {
		response.status = "Internal Server Error"
		cxt.IndentedJSON(cxt.StatusBadRequest,response)
		return
		
	}
	if input.password != input.comfirm_password{
		response.status = "password is mismatch"
		cxt.IndentedJSON(cxt.StatusBadRequest,response)
		return
	}

	for _,user:= range users{
		if user.UserName == input.user_name{
			response.status = "This user name already exist"
			cxt.IndentedJSON(cxt.StatusBadRequest,response)
			return
		}

		if user.Email == input.email{
			response.status = "This email already exist"
			cxt.IndentedJSON(cxt.StatusBadRequest,response)
			return
		}

	}

	usr:= entity.User{
		Name:input.name,
		LastName:input.last_name
		UserName: input.user_name
		Email: input.email
		password:input.password
	}
	
	user,ers:= usrHandler.userSrv.CreateUser(usr)

	if len(ers)> 0 {
		response.status = "Internal Server Error"
		cxt.IndentedJSON(cxt.StatusInternalServerError,response)
		return
	}

	response.status = "successfully created"
	response.user = user
	cxt.IndentedJSON(http.StatusCreated,response)
}

func(usrHandler *UserHandler) ChangeProfile(cxt * gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		user *entity.User
	}{
		status:"Unauthorized user"
	}
	input := &struct{
		name string
		last_name string		
		user_name string
		email string
		
	}{}

	sess := ursHandler.sessionHa.GetSession(r)
	if sess != nil{
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}

	if er = cxt.BindJSON(&input); er!=nil{
		response.status = "bad request..." 
		cxt.IndentedJSON(cxt.StatusBadRequest,response)
		return
	}

	if input.UserName =="" || input.Email == "" || input.Name=="" || input.LastName {
		response.status = "Invalid input..."
		cxt.IndentedJSON(cxt.StatusBadRequest,response)
		return
	}

	if user.UserName == input.user_name{
		response.status = "This user name already exist"
		cxt.IndentedJSON(cxt.StatusBadRequest,response)
		return
	}

	usr:= entity.User{
		Name:input.name,
		LastName:input.last_name
		UserName: input.user_name
		Email: input.email
	}
	usr.ID = sess.UserId
	user,ers:=usrHandler.userSrv.UpdateUser(usr)

	if ers!=nil {
		response.status = "Internal Server Error"
		cxt.IndentedJSON(cxt.StatusBadRequest,response)
		return
	}

	response.status = "profile successfully updated "
	response.user = user
	cxt.JSON(200,response)

}
func(usrHandler *UserHandler)ChangePassword(cxt * gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		
	}{
		status:"Unauthorized user"
	}
	input := &struct{
		old_password string
		new_password string
	}{}
	
	sess := ursHandler.sessionHa.GetSession(r)
	if sess != nil {
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}

	if er = cxt.BindJSON(&input); er!=nil{
		response.status = "bad request..." 
		cxt.IndentedJSON(cxt.StatusBadRequest,response)
		return
	}

	if input.old_password =="" || input.new_password == "" {
		response.status = "Invalid input..."
		cxt.IndentedJSON(cxt.StatusBadRequest,response)
		return
	}
	usr,er := usrHandler.userSrv.GetUser(sess.UserId)

	if ers!=nil {
		response.status = "No such user"
		cxt.IndentedJSON(cxt.StatusNotFound,response)
		return
	}

	if !helper.ComparePassword(usr.Password,input.old_password){
		response.status = "Incorrect password sign in again"
		cxt.IndentedJSON(cxt.StatusBadRequest,response)
		usrHandler.sessionHa.DeleteSession(cxt)
		return
	}

	response.status = "password succefully changed"
	cxt.IndentedJSON(cxt.StatusOK,response)

}

func(usrHandler *UserHandler) DeleteAccount(cxt * gin.Context){

	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		user *entity.User
		
	}{
		status:"Unauthorized user"
	}

	sess := ursHandler.sessionHa.GetSession(r)
	if sess != nil {
		cxt.IndentedJSON(http.StatusUnathorized,response)
		return
	}
	user,ers:= usrHandler.userSrv.GetUser(sess.UserId)

	if len(ers)>0 || user != nil{
		response.status = "No such user"
		cxt.IndentedJSON(http.StatusNotFound,response)
		return
	}
	
	user,ers:= usrHandler.userSrv.DeleteUser(uint(id))

	if len(ers)> 0 {
		response.status = "Internal Server Error user"
		cxt.IndentedJSON(http.StatusInternalServerError,response)
		return
	}

	response.status = "succefully Deleted."
	response.user = user
	cxt.IndentedJSON(200,response)
}
