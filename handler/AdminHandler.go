package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	helper "github.com/motikingo/ecommerceRESTAPI-Go/Helper"
	user "github.com/motikingo/ecommerceRESTAPI-Go/customer"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	// "github.com/motikingo/ecommerceRESTAPI-Go/item"
	// "github.com/motikingo/ecommerceRESTAPI-Go/cart"
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
		Status string
		Admin *entity.Customer
	}{
		Status:"Unauthorized user",
	}
	input := &struct{
		Name string
		Last_name string		
		User_name string
		Email string
		Password string
		Comfirm_password string


	}{}

	if er := cxt.BindJSON(&input); er!=nil{
		response.Status = "bad request..." 
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}


	if input.User_name =="" || input.Email == "" || input.Name=="" || input.Last_name == ""|| input.Password == "" || input.Comfirm_password == ""{
		response.Status = "Invalid input..."
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}	
	
	
	if input.Password != input.Comfirm_password{
		response.Status = "password is mismatch"
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if admHandler.AdminSrv.GetUserByUserName(input.User_name)!= nil{
		response.Status = "This user name already exist"
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if admHandler.AdminSrv.GetUserByEmail(input.User_name){
		response.Status = "This email already exist"
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	input.Password = helper.SecurePassword(input.Password)
	admin := entity.Customer{
		Name:input.Name,
		LastName:input.Last_name,
		UserName: input.User_name,
		Email: input.Email,
		Password:input.Password,
		Role:"Admin",
	}
	
	adm,ers:= admHandler.AdminSrv.CreateUser(admin)

	if len(ers)> 0 {
		response.Status = "Internal Server Error"
		cxt.IndentedJSON(http.StatusInternalServerError,response)
		return
	}

	response.Status = "successfully created"
	response.Admin = adm
	cxt.IndentedJSON(http.StatusCreated,response)
}

func(admHandler *AdminHandler) ChangeProfile(cxt * gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		Status string
		Admin *entity.Customer
	}{
		Status:"Unauthorized user",
	}
	input := &struct{
		Name string
		Last_name string		
		User_name string
		Email string
		
	}{}

	sess := admHandler.sessionHa.GetSession(cxt)
	if sess != nil || sess.Role != "Admin"{
		cxt.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	if er := cxt.BindJSON(&input); er!=nil{
		response.Status = "bad request..." 
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if input.User_name =="" || input.Email == "" || input.Name=="" || input.Last_name =="" {
		response.Status = "Invalid input..."
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	adm,ers:=admHandler.AdminSrv.GetUser(sess.UserId)

	if len(ers)>0 {
		response.Status = "No Admin found"
		cxt.IndentedJSON(http.StatusNotFound,response)
		return
	}

	admin:= entity.Customer{
		Name:input.Name,
		LastName:input.Last_name,
		UserName: input.User_name,
		Email: input.Email,
	}
	admin.ID = adm.ID
	adm,ers = admHandler.AdminSrv.UpdateUser(admin)

	if ers!=nil {
		response.Status = "Internal Server Error"
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	response.Status = "profile successfully updated "
	response.Admin = adm
	cxt.JSON(200,response)
}

func(admHandler *AdminHandler)ChangePassword(cxt * gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		Status string
		
	}{
		Status:"Unauthorized user",
	}
	input := &struct{
		Old_password string
		New_password string
	}{}
	
	sess := admHandler.sessionHa.GetSession(cxt)
	if sess != nil || sess.Role != "Admin" {
		cxt.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	if er := cxt.BindJSON(&input); er!=nil{
		response.Status = "bad request..." 
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if input.Old_password =="" || input.New_password == "" {
		response.Status = "Invalid input..."
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}
	adm,er := admHandler.AdminSrv.GetUser(sess.UserId)

	if len(er)>0 {
		response.Status = "No such Admin"
		cxt.IndentedJSON(http .StatusNotFound,response)
		return
	}

	if !helper.ComparePassword(adm.Password,input.Old_password){
		response.Status = "Incorrect password sign in again"
		cxt.IndentedJSON(http.StatusBadRequest,response)
		admHandler.sessionHa.DeleteSession(cxt)
		return
	}

	response.Status = "password succefully changed"
	cxt.IndentedJSON(http.StatusOK,response)

}

func(admHandler *AdminHandler) AdminLogIn (cxt *gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		Status string
		
	}{
		Status:"Unauthorized user",
	}
	input := &struct{
		User_name string
		Password string
	}{}
	
	if er := cxt.BindJSON(&input); er!=nil{
		response.Status = "bad request..." 
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	adm := admHandler.AdminSrv.GetUserByUserName(input.User_name)

	if adm == nil{
		response.Status = "No Admin with this user name..."
		cxt.IndentedJSON(http.StatusNotFound,response)
		return
	}

	if !helper.ComparePassword(adm.Password,input.Password){
		response.Status = "Invalid password..."
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	session := &entity.Session{
		UserId:adm.ID,
		UserName:adm.UserName,
		Email:adm.Email,
		Role: "Admin",
	}

	if !admHandler.sessionHa.CreateSession(session,cxt){
		response.Status = "Internal server Error..."
		cxt.IndentedJSON(http.StatusInternalServerError,response)
		return
	}

	response.Status = "Successfully loged In..."
	cxt.IndentedJSON(http.StatusOK,response)
}


func(admHandler *AdminHandler)AdminLogOut(cxt *gin.Context){
	cxt.Header("Content-Type","application/json")
	response := &struct{
		Status string
		
	}{
		Status:"Unauthorized user",
	}
	fmt.Println("here")
	sess := admHandler.sessionHa.GetSession(cxt)
	fmt.Println(sess)
	if sess == nil {
		fmt.Println("here")
		cxt.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	adm := admHandler.AdminSrv.GetUserByUserName(sess.UserName)

	if adm == nil{
		response.Status = "No user with this user name..."
		cxt.IndentedJSON(http.StatusNotFound,response)
		return
	}

	if !admHandler.sessionHa.DeleteSession(cxt){
		response.Status = "Internal server Error..."
		cxt.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	response.Status = "Successfully loged out..."
	cxt.IndentedJSON(http.StatusOK,response)
}

