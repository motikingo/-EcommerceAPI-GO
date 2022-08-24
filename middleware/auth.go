package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/motikingo/ecommerceRESTAPI-Go/handler"
)

type Middleware struct{
	session *handler.SessionHandler
}

func NewMiddlerware(session *handler.SessionHandler) *Middleware{
	return &Middleware{session: session}
}

func (mid Middleware) UserLogedIn()gin.HandlerFunc{
	return gin.HandlerFunc(func(ctx *gin.Context){
		session := mid.session.GetSession(ctx)

		if session == nil{
			ctx.AbortWithStatusJSON(401,gin.H{"status":"User already loged in"})
			return 
		}
		ctx.Next()
	})
}

func (mid Middleware) AdminLogedIn()gin.HandlerFunc{
	return gin.HandlerFunc(func(ctx *gin.Context){
		session := mid.session.GetSession(ctx)

		if session == nil || session.Role == "Admin"{
			ctx.AbortWithStatusJSON(401,gin.H{"status":"User already loged in"})
			return 
		}
		ctx.Next()
	})
}
