package entity

import (
	"time"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
)

type Customer struct{
	gorm.Model  
	Name string  `json:"name"`
	LastName string `json:"lastName"`
	UserName string	`json:"username"`
	Email string	`json:"email"`
	Password string	 `json:"password"`
	
}

type Catagory struct{
	gorm.Model
	Name string `json:"name"`
	Image string 	`json:"image"`
	Description string `json:"description"`
	Items_Id []uint		`json:"items_Id"`
}

type Item struct{
	gorm.Model
	Name string 			`json:"name"`
	Description string		`json :"description"`
	Brand string 			`json:"brand"`
	Image string 			`json:"image"`
	Price uint 				`json:"price"`
	Number int				 `json:"number"`
	ProductionDate time.Time `json:"production_date"`
	ExpireDate time.Time 	`json:"expire_date"`

}

type Cart struct{
	gorm.Model
	UserId uint	`json:user_Id`
	Items []uint 	`json:"items"`
}


type Session struct{
	UserId uint
	UserName string
	Email string
	Role string
	jwt.StandardClaims

}

