package handler

import (

	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	"github.com/motikingo/ecommerceRESTAPI-Go/customer"
	// "github.com/motikingo/ecommerceRESTAPI-Go/item"
	// "github.com/motikingo/ecommerceRESTAPI-Go/cart"
	"github.com/motikingo/ecommerceRESTAPI-Go/Helper"

)

type AdminHandler struct {

	AdminSrv user.UserService
	sessionHa *SessionHandler
}

func NewAdminHandler(AdminSrv user.UserService, sessionHa *SessionHandler)AdminHandler{
	return AdminHandler{AdminSrv: AdminSrv,sessionHa:sessionHa}
}

func (admHandler *AdminHandler)GetAdmins(cxt *gin.Context){
	cxt.Header("Content-Type","application/json")
	sess := admHandler.sessionHa.GetSession(cxt)
	if sess != nil || sess.Role != "Admin"{
		cxt.IndentedJSON(http.StatusUnauthorized,gin.H{"Error":"Unauthorized user"})
		return
	}
	adms,e:= admHandler.AdminSrv.GetUsers()

	if len(e)>0 {
		cxt.IndentedJSON(http.StatusInternalServerError,"Internal Server Error")
		return
	}
	if adms == nil{
		cxt.IndentedJSON(http.StatusOK,"No Admin Found")
		return
	}

	cxt.IndentedJSON(200,adms)

}

func (admHandler *AdminHandler)GetAdmin(cxt *gin.Context){
	cxt.Header("Content-Type","application/json")
	sess := admHandler.sessionHa.GetSession(cxt)
	if sess != nil || sess.Role != "Admin"{
		cxt.IndentedJSON(http.StatusUnauthorized,"Unauthorized user")
		return
	}

	adm,e:= admHandler.AdminSrv.GetUser(sess.UserId)

	if len(e)>0{
		cxt.IndentedJSON(http.StatusInternalServerError,"Internal Server Error")
		return

	}

	cxt.JSON(200,adm)

}

func(admHandler *AdminHandler) CreateAdmin(cxt *gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		Admin *entity.Customer
	}{
		status:"Unauthorized user",
	}
	input := &struct{
		name string
		last_name string		
		user_name string
		email string
		password string
		comfirm_password string


	}{}

	if er := cxt.BindJSON(&input); er!=nil{
		response.status = "bad request..." 
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	input.password = helper.SecurePassword(input.password)

	if input.user_name =="" || input.email == "" || input.name=="" || input.last_name == ""|| input.password == "" || input.comfirm_password == ""{
		response.status = "Invalid input..."
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}	
	
	
	if input.password != input.comfirm_password{
		response.status = "password is mismatch"
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if admHandler.AdminSrv.GetUserByUserName(input.user_name)!= nil{
		response.status = "This user name already exist"
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if admHandler.AdminSrv.GetUserByEmail(input.user_name){
		response.status = "This email already exist"
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	admin := entity.Customer{
		Name:input.name,
		LastName:input.last_name,
		UserName: input.user_name,
		Email: input.email,
		Password:input.password,
	}
	
	adm,ers:= admHandler.AdminSrv.CreateUser(admin)

	if len(ers)> 0 {
		response.status = "Internal Server Error"
		cxt.IndentedJSON(http.StatusInternalServerError,response)
		return
	}

	response.status = "successfully created"
	response.Admin = adm
	cxt.IndentedJSON(http.StatusCreated,response)
}

func(admHandler *AdminHandler) ChangeProfile(cxt * gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		Admin *entity.Customer
	}{
		status:"Unauthorized user",
	}
	input := &struct{
		name string
		last_name string		
		user_name string
		email string
		
	}{}

	sess := admHandler.sessionHa.GetSession(cxt)
	if sess != nil || sess.Role != "Admin"{
		cxt.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	if er := cxt.BindJSON(&input); er!=nil{
		response.status = "bad request..." 
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if input.user_name =="" || input.email == "" || input.name=="" || input.last_name =="" {
		response.status = "Invalid input..."
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	adm,ers:=admHandler.AdminSrv.GetUser(sess.UserId)

	if len(ers)>0 {
		response.status = "No Admin found"
		cxt.IndentedJSON(http.StatusNotFound,response)
		return
	}

	admin:= entity.Customer{
		Name:input.name,
		LastName:input.last_name,
		UserName: input.user_name,
		Email: input.email,
	}
	admin.ID = sess.UserId
	adm,ers = admHandler.AdminSrv.UpdateUser(admin)

	if ers!=nil {
		response.status = "Internal Server Error"
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	response.status = "profile successfully updated "
	response.Admin = adm
	cxt.JSON(200,response)
}

func(admHandler *AdminHandler)ChangePassword(cxt * gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		
	}{
		status:"Unauthorized user",
	}
	input := &struct{
		old_password string
		new_password string
	}{}
	
	sess := admHandler.sessionHa.GetSession(cxt)
	if sess != nil || sess.Role != "Admin" {
		cxt.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	if er := cxt.BindJSON(&input); er!=nil{
		response.status = "bad request..." 
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if input.old_password =="" || input.new_password == "" {
		response.status = "Invalid input..."
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}
	adm,er := admHandler.AdminSrv.GetUser(sess.UserId)

	if len(er)>0 {
		response.status = "No such Admin"
		cxt.IndentedJSON(http .StatusNotFound,response)
		return
	}

	if !helper.ComparePassword(adm.Password,input.old_password){
		response.status = "Incorrect password sign in again"
		cxt.IndentedJSON(http.StatusBadRequest,response)
		admHandler.sessionHa.DeleteSession(cxt)
		return
	}

	response.status = "password succefully changed"
	cxt.IndentedJSON(http.StatusOK,response)

}

func(admHandler *UserHandler) AdminLogIn (cxt * gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		
	}{
		status:"Unauthorized user",
	}
	input := &struct{
		user_name string
		password string
	}{}
	
	if er := cxt.BindJSON(&input); er!=nil{
		response.status = "bad request..." 
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	adm := admHandler.userSrv.GetUserByUserName(input.user_name)

	if adm == nil{
		response.status = "No Admin with this user name..."
		cxt.IndentedJSON(http.StatusNotFound,response)
		return
	}

	if !helper.ComparePassword(adm.Password,input.password){
		response.status = "Invalid password..."
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	session := entity.Session{
		UserId:adm.ID,
		UserName:adm.UserName,
		Email:adm.Email,
		Role: "Admin",
	}

	if !admHandler.sessionHa.CreateSession(session,cxt){
		response.status = "Internal server Error..."
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	response.status = "Successfully loged In..."
	cxt.IndentedJSON(http.StatusOK,response)
}


func(admHandler *UserHandler)AdminLogout(cxt * gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		status string
		
	}{
		status:"Unauthorized user",
	}

	sess := admHandler.sessionHa.GetSession(cxt)
	if sess != nil {
		cxt.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	adm := admHandler.userSrv.GetUserByUserName(sess.Email)

	if adm == nil{
		response.status = "No user with this user name..."
		cxt.IndentedJSON(http.StatusNotFound,response)
		return
	}

	if !admHandler.sessionHa.DeleteSession(cxt){
		response.status = "Internal server Error..."
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	response.status = "Successfully loged out..."
	cxt.IndentedJSON(http.StatusOK,response)
}

