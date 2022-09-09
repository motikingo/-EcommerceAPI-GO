package handler

import (
	//"fmt"
	//"net/http"
	//"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	helper "github.com/motikingo/ecommerceRESTAPI-Go/Helper"
	user "github.com/motikingo/ecommerceRESTAPI-Go/customer"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	"github.com/motikingo/ecommerceRESTAPI-Go/item"
	"github.com/motikingo/ecommerceRESTAPI-Go/record"
)

type UserHandler struct {

	userSrv user.UserService
	itemSrv item.ItemService
	cartHa 	*CartHandler
	RecordSrv record.RecordService
	sessionHa *SessionHandler
}

func NewUserHandler(userSrv user.UserService,itemSrv item.ItemService, cartHa *CartHandler,RecordSrv record.RecordService,sessionHa *SessionHandler)UserHandler{
	return UserHandler{userSrv: userSrv,itemSrv:itemSrv,cartHa:cartHa,RecordSrv:RecordSrv,sessionHa:sessionHa}
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
		Name string
		Last_name string		
		User_name string
		Email string
		Password string
		Comfirm_password string


	}{}

	if er := ctx.BindJSON(&input); er!=nil{
		response.status = "bad request..." 
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if input.User_name =="" || input.Email == "" || input.Name=="" || input.Last_name=="" || input.Password == "" || input.Comfirm_password == ""{
		response.status = "Invalid input..."
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if input.Password != input.Comfirm_password{
		response.status = "password is mismatch"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if usrHandler.userSrv.GetUserByUserName(input.User_name) != nil {
		response.status = "This user name already exist"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if usrHandler.userSrv.GetUserByEmail(input.User_name){
		response.status = "This email already exist"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}
	input.Password = helper.SecurePassword(input.Password)
	usr := entity.Customer{
		Name:input.Name,
		LastName:input.Last_name,
		UserName: input.User_name,
		Email: input.Email,
		Password:input.Password,
	}
	
	user,ers:= usrHandler.userSrv.CreateUser(usr)

	if len(ers)> 0 {
		response.status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	fmt.Println("hereeeee")
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
		Status string
		
	}{
		Status:"Unauthorized user",
	}
	input := &struct{
		User_name string
		Password string
	}{}
	
	if er := ctx.BindJSON(&input); er!=nil{
		response.Status = "bad request..." 
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	user := usrHandler.userSrv.GetUserByUserName(input.User_name)

	if user == nil{
		response.Status = "No user with this user name..."
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}

	if !helper.ComparePassword(user.Password,input.Password){
		response.Status = "Invalid password..."
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	session := &entity.Session{
		UserId:user.ID,
		UserName:user.UserName,
		Email:user.Email,
	}
	
	if !usrHandler.sessionHa.CreateSession(session,ctx){
		fmt.Println("here123")
		response.Status = "Internal server Error..."
		ctx.IndentedJSON(http.StatusInternalServerError,session)
		return
	}
	response.Status = "Successfully loged In..."
	ctx.IndentedJSON(http.StatusOK,response)
}


func(usrHandler *UserHandler)Logout(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		Status string
		
	}{
		Status:"Unauthorized user",
	}

	sess := usrHandler.sessionHa.GetSession(ctx)
	if sess == nil {
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	user := usrHandler.userSrv.GetUserByUserName(sess.UserName)

	if user == nil{
		response.Status = "No user with this user name..."
		ctx.IndentedJSON(http.StatusNotFound,response)
		return
	}
	if !usrHandler.sessionHa.DeleteSession(ctx){
		response.Status = "Internal server Error..."
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	response.Status = "Successfully Loged out..."
	ctx.IndentedJSON(http.StatusOK,response)
}

func(usrHandler *UserHandler) AddItemToMyCart(ctx * gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		Status string
		Cart entity.Cart
		
	}{
		Status:"Unauthorized user",
	}
	input := &struct{
		Item_Id uint
		HowMany int
	}{}
	
	sess := usrHandler.sessionHa.GetSession(ctx)
	if sess == nil {
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}	

	if er := ctx.BindJSON(&input); er!=nil{
		response.Status = "bad request..." 
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if input.Item_Id == 0 || input.HowMany<1{
		response.Status = "Invalid Input..." 
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	} 
	item,ers := usrHandler.itemSrv.GetItem(input.Item_Id)
	
	if len(ers)>0{
		response.Status = "No such Item exist"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}

	if item.Number < input.HowMany{
		response.Status = "Appology we run out of Item for now"
		ctx.IndentedJSON(http.StatusOK,response)
		return
	}

	cart := usrHandler.cartHa.GetCart(ctx)

	if cart == nil{
		response.Status = "Add cart first please..." 
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}
	
	cart.Items = append(cart.Items, entity.ItemInfo{
		ItemId : input.Item_Id,
		ItemName:item.Name,
		Number:  input.HowMany,
		ItemBill:float64(input.HowMany) * item.Price,
	})

	if ! usrHandler.cartHa.UpdateCart(*cart,ctx) {
		response.Status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}

	item.Number  = item.Number - input.HowMany
	_,ers = usrHandler.itemSrv.UpdateItem(*item)

	if len(ers)>0{
		cart.Items = cart.Items[:len(cart.Items)-1]
		response.Status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	response.Status = "Item Added to cart successfully..."
	response.Cart = *cart 
	ctx.IndentedJSON(http.StatusBadRequest,response)
}

func(usrHandler *UserHandler) DeleteItemFromMyCart(ctx * gin.Context){

	ctx.Header("Content-Type","application/json")
	response := &struct{
		Status string
		
	}{
		Status:"Unauthorized user",
	}
	
	
	sess := usrHandler.sessionHa.GetSession(ctx)
	if sess == nil {
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}	

	id,_ := strconv.Atoi(ctx.Param("item_id")) 

	item,ers := usrHandler.itemSrv.GetItem(uint(id))
	
	if len(ers)>0{
		response.Status = "No such Item exist"
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}
	
	cart := usrHandler.cartHa.GetCart(ctx)

	if cart == nil{
		response.Status = "Add cart first please..." 
		ctx.IndentedJSON(http.StatusBadRequest,response)
		return
	}
	var items []entity.ItemInfo
	for _,itmInfo:= range cart.Items{
		if itmInfo.ItemId != item.ID{
			items = append(items, itmInfo)
		}
	}
	cart.Items = items 
	if !usrHandler.cartHa.UpdateCart(*cart,ctx){
		response.Status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	item.Number  = item.Number + cart.Items[item.ID].Number
	_,ers = usrHandler.itemSrv.UpdateItem(*item)

	if len(ers)>0{
		//cart.Items = append(cart.Items,)
		response.Status = "Internal Server Error"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	response.Status = "Successfully remote Item from cart"
	ctx.IndentedJSON(http.StatusInternalServerError,response)

}

func(usrHandler *UserHandler) DeleteMyCart(ctx * gin.Context){

	cart:= usrHandler.cartHa.DeleteCart(ctx)
	for _, info := range cart.Items{
		item,er:= usrHandler.itemSrv.GetItem(info.ItemId)
		if len(er)>0{
			ctx.IndentedJSON(http.StatusInternalServerError,gin.H{"status" : "No such Item"})
			return
		}
		item.Number += info.Number
		_,er = usrHandler.itemSrv.UpdateItem(*item)

		if len(er)>0{
			ctx.IndentedJSON(http.StatusInternalServerError,gin.H{"status" : "Internal Server Error"})
			return
		}
	}

}

func (usrHandler *UserHandler) Order(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	response := &struct{
		Status string
		TotalBill float64
		Record entity.Record
		
	}{
		Status:"Unauthorized user",
	}

	sess := usrHandler.sessionHa.GetSession(ctx)
	if sess == nil {
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}

	cart:= usrHandler.cartHa.GetCart(ctx)

	if cart == nil{
		response.Status = "There is no cart with this user Id"
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	if len(cart.Items) == 0{
		response.Status = "cart is empty Add item first please..."
		ctx.IndentedJSON(http.StatusUnauthorized,response)
		return
	}
	for _,bill := range cart.Items{
		
		response.TotalBill += bill.ItemBill
	} 
	record := usrHandler.RecordSrv.GetRecordByUserID(sess.UserId)

	if record ==nil{
		reco := entity.Record{
			AddedAt: time.Now(),
			UserId: sess.UserId,
			Cart_Infos: make([]entity.CartInfo,0),
		}
		
		if record = usrHandler.RecordSrv.CreateRecord(reco); record == nil{
			response.Status = "Internal Server Error"
			ctx.IndentedJSON(http.StatusInternalServerError,response)
			return
		}
		
	}
	cartInfo := &entity.CartInfo{
		RecordUserId:record.UserId,
		Item_Infos: make([]entity.ItemInfo, 0),
	}
	cartInfo = usrHandler.RecordSrv.CreateInfo(*cartInfo)

	var items []entity.ItemInfo

	for _,itm := range cart.Items{
		//var item entity.ItemInfo
		item := itm
		item.CartInfoID = cartInfo.ID
		//fmt.Println(cartInfo.ID)
		items = append(items, item)
	}
	cartInfo.Item_Infos  = items
	
	record.Cart_Infos = append(record.Cart_Infos,*cartInfo) 
	reco := usrHandler.RecordSrv.UpdateRecord(*record)
	// fmt.Println("we here")
	if reco == nil {
		response.Status = "Internal Server Error user"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	cart.Items = []entity.ItemInfo{}
	if !usrHandler.cartHa.UpdateCart(*cart,ctx){
		response.Status = "can't Empty Cart..."
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	}
	response.Status = "total bill calculated..."
	response.Record = *record
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

	_ = usrHandler.cartHa.DeleteCart(ctx)
	r := usrHandler.RecordSrv.GetRecordByUserID(user.ID)
	r= usrHandler.RecordSrv.ClearRecord(r.UserId)
	
	if !usrHandler.sessionHa.DeleteSession(ctx) || r==nil{
		response.status = "Internal Server Error user"
		ctx.IndentedJSON(http.StatusInternalServerError,response)
		return
	} 


	response.status = "succefully Deleted."
	response.user = user
	ctx.IndentedJSON(200,response)
}
