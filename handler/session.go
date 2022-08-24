package handler

import (
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

func (sessionHa *SessionHandler) CreateSession(session entity.Session, ctx *gin.Context)bool{
	ctx.Header("Content-Type","application/json")

	expireTime :=time.Now().Add(24 * time.Hour)
		 
	session.StandardClaims = jwt.StandardClaims{
		ExpiresAt : expireTime.Unix(),		
	}

	tkn:= jwt.NewWithClaims(jwt.SigningMethodHS256,session)

	tknstr,er := tkn.SignedString(hash_key)

	if er!=nil{
		ctx.JSON(404,tkn.Claims.Valid().Error())
		return false
	}

	ctx.SetCookie(tknstr,sessionName,int(expireTime.Unix()),"/","",true,true)
	return true
}

func (sessionHa *SessionHandler)DeleteSession(ctx *gin.Context)bool {
	ctx.Header("Content-Type","application/json")
	session := entity.Session{}
	expireTime :=time.Now().Add(-24 * time.Hour)
		 
	session.StandardClaims = jwt.StandardClaims{
		ExpiresAt : expireTime.Unix(),	
	}

	tkn:= jwt.NewWithClaims(jwt.SigningMethodHS256,session)

	tknstr,er := tkn.SignedString(hash_key)

	if er!=nil{
		ctx.IndentedJSON(404,tkn.Claims.Valid().Error())
		return false
	}
	ctx.SetCookie(sessionName,tknstr,int(expireTime.Unix()),"/","",true,true)
	return true
}

func(sessionHa *SessionHandler)GetSession(ctx *gin.Context) *entity.Session{
	ctx.Header("Content-Type","application/json")
	cookie,er := ctx.Cookie(sessionName)
	var session *entity.Session
	if er!=nil || cookie == "" {
		return nil
	}
	tkn,err := jwt.ParseWithClaims(cookie,session,func(t *jwt.Token) (interface{}, error) {
		return hash_key,nil
	})

	if err != nil || !tkn.Valid{
		return nil
	}

	return session
}