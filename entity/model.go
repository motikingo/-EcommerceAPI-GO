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
	Role 	string 	`json:"role"`
}

type Catagory struct{
	gorm.Model
	Name string `json:"name"`
	Image string 	`json:"image"`
	Description string `json:"description"`
	Items []Item		`json:"items" gorm:"Many2Many:catagory_items"`
}

type Item struct{
	gorm.Model	
	Name string 			`json:"name"`
	Description string		`json :"description"`
	Brand string 			`json:"brand"`
	Image string 			`json:"image"`
	Price float64 			`json:"price"`
	Number int				`json:"number"`
	ProductionDate time.Time`json:"production_date"`
	ExpireDate time.Time	`json:"expire_date"`
	Catagories []Catagory 	`json:"catagories"  gorm:"Many2Many:catagory_items"`
}

type Record struct{
	AddedAt time.Time
	UserId uint `gorm:"Primary_Key:user_id"`
	Cart_Infos []CartInfo
}
type CartInfo struct{
	gorm.Model
	RecordUserId uint
	Item_Infos []ItemInfo	 `json:"items"`
}

type Cart struct{
	UserId uint			 `json:user_Id`
	Items []ItemInfo	 `json:"items"`
	jwt.RegisteredClaims
}
type ItemInfo struct{
	ID uint 	`gorm:"Primary_Key:id; AUTO_INCREMENT"`
	ItemId uint 
	ItemName string 
	Number int
	ItemBill float64
	CartInfoID uint
}
type Session struct{
	UserId uint
	UserName string
	Email string
	Role string
	jwt.StandardClaims
}

