package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
)


const(
	sessionName = "ecommerce"
	hash_key = "hash key"
)

type SessionHandler struct {
}

func NewSessionHandler() SessionHandler {
	return SessionHandler{}
}

func (sessionHa *SessionHandler) CreateSession(session entity.Session, ctx *gin.Context)string{
	ctx.Header("Content-Type","application/json")

	expireTime :=time.Now().Add(24 * time.Hour)
		 
	sess := session.StandardClaims(
		ExpireAt : expireTime.Unix()
		
	)

	tkn:= jwt.NewWithClaims(jwt.SigningMethodHS256,sess)

	tknstr,er := tkn.SignedString(hash_key)

	if er!=nil{
		ct.JSON(404,tkn.Claims.Valid().Error())
		return
	}

	cookie := *http.Cookie{
		Name: sessionName,
		Value: tknstr,
		Expires: expireTime,
	}

	http.SetCookie(w,cookie)

}

func (sessionHa *SessionHandler)DeleteSession(ctx *gin.Context){
	ctx.Header("Content-Type","application/json")
	session := entity.Session{}
	expireTime :=time.Now().Add(24 * time.Hour)
		 
	sess := session.StandardClaims(
		ExpireAt : expireTime.Unix()
		
	)

	tkn:= jwt.NewWithClaims(jwt.SigningMethodHS256,sess)

	tknstr,er := tkn.SignedString(hash_key)

	if er!=nil{
		ct.JSON(404,tkn.Claims.Valid().Error())
		return
	}

	ctx.SetCookie(sessionName,tknstr,60*60*24,"/","",true,true)
	
}

func(sessionHa *SessionHandler)GetSession(ctx *gin.Context) *entity.Session{
	ctx.Header("Content-Type","application/json")
	cookie,er := ctx.Cookie(sessionName)

	if er || cookie == nil{
		return nil
	}
	tkn,err := jwt.ParseWithClaims(cookie.Value,jwt.StandardClaims,func(t *jwt.Token) (interface{}, error) {
		return hash_key
	})

	if err != nil || !tkn.Valid{
		return nil
	}

	session := tkn.Claims


	return session
}