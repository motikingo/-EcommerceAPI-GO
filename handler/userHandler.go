package handler

import (
	//"fmt"
	//"net/http"
	//"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/motikingo/ecommerceRESTAPI-Go/cart"
	"github.com/motikingo/ecommerceRESTAPI-Go/customer"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	"github.com/motikingo/ecommerceRESTAPI-Go/item"
	"github.com/motikingo/ecommerceRESTAPI-Go/Helper"

)

type UserHandler struct {

	userSrv user.UserService
	itemSrv item.ItemService
	cartsrv cart.CartService
	sessionHa *SessionHandler
}

func NewUserHandler(userSrv user.UserService, sessionHa *SessionHandler)UserHandler{
	return UserHandler{userSrv: userSrv,sessionHa:sessionHa}
}

func (usrHandler *UserHandler)GetUsers(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	sess := usrHandler.sessionHa.GetSession(ctx)
	if sess != nil || sess.Role != "Admin"{
		ctx.IndentedJSON(http.StatusUnauthorized,"Unauthorized user")
		return
	}
	users,e:= usrHandler.userSrv.GetUsers()

	if len(e)>0 {
		ctx.IndentedJSON(http.StatusInternalServerError,"Internal Server Error")
		return
	}
	if users == nil{
		ctx.IndentedJSON(http.StatusOK,"No user Found")
		return
	}

	ctx.IndentedJSON(200,users)

}

func (usrHandler *UserHandler)GetUser(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	sess := usrHandler.sessionHa.GetSession(ctx)
	if sess != nil || sess.Role != "Admin"{
		ctx.IndentedJSON(http.StatusUnauthorized,"Unauthorized user")
		return
	}

	user,e:= usrHandler.userSrv.GetUser(sess.UserId)

	if len(e)>0{
		ctx.IndentedJSON(http.StatusInternalServerError,"Internal Server Error")
		return

	}

	ctx.JSON(200,user)

}

func(usrHandler *UserHandler) CreateUser(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		user *entity.Customer
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

	if er := ctx.BindJSON(&input); er!=nil{
		response.status = "bad request..." 
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	input.password = helper.SecurePassword(input.password)

	if input.user_name =="" || input.email == "" || input.name=="" || input.last_name=="" || input.password == "" || input.comfirm_password == ""{
		response.status = "Invalid input..."
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if input.password != input.comfirm_password{
		response.status = "password is mismatch"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if usrHandler.userSrv.GetUserByUserName(input.user_name) != nil {
		response.status = "This user name already exist"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if usrHandler.userSrv.GetUserByEmail(input.user_name){
		response.status = "This email already exist"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	usr := entity.Customer{
		Name:input.name,
		LastName:input.last_name,
		UserName: input.user_name,
		Email: input.email,
		Password:input.password,
	}
	
	user,ers:= usrHandler.userSrv.CreateUser(usr)

	if len(ers)> 0 {
		response.status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}

	response.status = "successfully created"
	response.user = user
	ctx.IndentedJSON(http.StatusCreated,response)
}

func(usrHandler *UserHandler) ChangeProfile(ctx * gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		user *entity.Customer
	}{
		status:"Unauthorized user",
	}
	input := &struct{
		name string
		last_name string		
		user_name string
		email string
		
	}{}

	sess := usrHandler.sessionHa.GetSession(ctx)
	if sess != nil{
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	if er := ctx.BindJSON(&input); er!=nil{
		response.status = "bad request..." 
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if input.user_name =="" || input.email == "" || input.name=="" || input.last_name=="" {
		response.status = "Invalid input..."
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	user,ers:=usrHandler.userSrv.GetUser(sess.UserId)

	if ers!=nil {
		response.status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if user.UserName == input.user_name{
		response.status = "This user name already exist"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	usr:= entity.Customer{
		Name:input.name,
		LastName:input.last_name,
		UserName: input.user_name,
		Email: input.email,
	}
	usr.ID = sess.UserId
	user,ers = usrHandler.userSrv.UpdateUser(usr)

	if ers!=nil {
		response.status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	response.status = "profile successfully updated "
	response.user = user
	ctx.JSON(200,response)

}
func(usrHandler *UserHandler)ChangePassword(ctx * gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		
	}{
		status:"Unauthorized user",
	}
	input := &struct{
		old_password string
		new_password string
	}{}
	
	sess := usrHandler.sessionHa.GetSession(ctx)
	if sess != nil {
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	if er := ctx.BindJSON(&input); er!=nil{
		response.status = "bad request..." 
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if input.old_password =="" || input.new_password == "" {
		response.status = "Invalid input..."
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}
	usr,er := usrHandler.userSrv.GetUser(sess.UserId)

	if er!=nil {
		response.status = "No such user"
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}

	if !helper.ComparePassword(usr.Password,input.old_password){
		response.status = "Incorrect password sign in again"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		usrHandler.sessionHa.DeleteSession(ctx)
		return
	}

	response.status = "password succefully changed"
	ctx.IndentedJSON(http.StatusOK,response)

}

func(usrHandler *UserHandler)LogIn(ctx * gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		
	}{
		status:"Unauthorized user",
	}
	input := &struct{
		user_name string
		password string
	}{}
	
	if er := ctx.BindJSON(&input); er!=nil{
		response.status = "bad request..." 
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	user := usrHandler.userSrv.GetUserByUserName(input.user_name)

	if user == nil{
		response.status = "No user with this user name..."
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}

	if !helper.ComparePassword(user.Password,input.password){
		response.status = "Invalid password..."
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	session := entity.Session{
		UserId:user.ID,
		UserName:user.UserName,
		Email:user.Email,
	}
	if !usrHandler.sessionHa.CreateSession(session,ctx){
		response.status = "Internal server Error..."
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	response.status = "Successfully loged In..."
	ctx.IndentedJSON(http.StatusOK,response)
}


func(usrHandler *UserHandler)Logout(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		
	}{
		status:"Unauthorized user",
	}

	sess := usrHandler.sessionHa.GetSession(ctx)
	if sess != nil {
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	user := usrHandler.userSrv.GetUserByUserName(sess.Email)

	if user == nil{
		response.status = "No user with this user name..."
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}

	if !usrHandler.sessionHa.DeleteSession(ctx){
		response.status = "Internal server Error..."
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	response.status = "Successfully loged out..."
	ctx.IndentedJSON(http.StatusOK,response)
}

func(usrHandler *UserHandler) AddItemToMyCart(ctx * gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		
	}{
		status:"Unauthorized user",
	}
	input := &struct{
		item_Id uint
		howMany int
	}{}
	
	sess := usrHandler.sessionHa.GetSession(ctx)
	if sess != nil {
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}	

	if er := ctx.BindJSON(&input); er!=nil{
		response.status = "bad request..." 
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if string(input.item_Id) == "" || input.howMany<1{
		response.status = "Invalid Input..." 
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	} 
	item,ers := usrHandler.itemSrv.GetItem(input.item_Id)
	
	if len(ers)>0{
		response.status = "No such Item exist"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if item.Number < input.howMany{
		response.status = "Appology we run out of Item for now"
		ctx.IndentedJSON(http.StatusOK,response)
		return
	}

	cart := usrHandler.cartsrv.GetCartByUserID(sess.UserId)

	if cart == nil{
		response.status = "Add cart first please..." 
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}
	cart.Items[input.item_Id] = float64(input.howMany) * item.Price

	cart,ers = usrHandler.cartsrv.UpdateCart(*cart)

	if len(ers)>0{
		response.status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}

	item.Number  = item.Number - input.howMany
	item,ers = usrHandler.itemSrv.UpdateItem(*item)

	if len(ers)>0{
		delete(cart.Items,item.ID)
		response.status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
}

func (usrHandler *UserHandler) Order(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		totalBill float64
		
	}{
		status:"Unauthorized user",
	}

	sess := usrHandler.sessionHa.GetSession(ctx)
	if sess != nil {
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	cart:= usrHandler.cartsrv.GetCartByUserID(sess.UserId)

	if cart == nil{
		response.status = "There is no cart with this user Id"
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	for _,bill := range cart.Items{
		
		response.totalBill += bill
	} 

	response.status = "total bill calculated..."
	ctx.IndentedJSON(200,response)

}

func(usrHandler *UserHandler) DeleteAccount(ctx * gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		status string
		user *entity.Customer
		
	}{
		status:"Unauthorized user",
	}
	sess := usrHandler.sessionHa.GetSession(ctx)
	if sess != nil {
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	user,ers:= usrHandler.userSrv.GetUser(sess.UserId)

	if len(ers)>0 || user != nil{
		response.status = "No such user"
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}
	
	user,ers = usrHandler.userSrv.DeleteUser(sess.UserId)

	if len(ers)> 0 {
		response.status = "Internal Server Error user"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}

	response.status = "succefully Deleted."
	response.user = user
	ctx.IndentedJSON(200,response)
}
