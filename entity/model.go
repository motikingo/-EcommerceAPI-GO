package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct{
	gorm.Model  
	Name string  `json:"name"`
	LastName string `json:"lastName"`
	UserName string	`json:"username"`
	Email string	`json:"email"`
	Password string	 `json:"password"`
	CartId uint 	`json:"cart_id"`
	
}

type Catagory struct{
	gorm.Model
	Name string `json:"name"`
	Description string `json:"description"`
	Items []Item 		`json:"items"`

}

type Item struct{
	gorm.Model
	Name string 	`json:"name"`
	Catagorys []Catagory `json:"catagories"`
	Image string 		`json:"image"`
	Price uint 			`json:"price"`
	Description *Description `json:"description"`

}

type Description struct{
	gorm.Model
	Brand string 	`json:"brand"`
	ProductionDate time.Time `json:"production_date"`
	ExpireDate time.Time `json:"expire_date"`
}

type Cart struct{
	gorm.Model
	CustomerId uint `json:"customer_id"`
	Items []Item 	`json:"items"`
}

