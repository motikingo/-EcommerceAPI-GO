package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
)

func Connect() *gorm.DB {

	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal(err)
		return nil
	}

	dialect := os.Getenv("dialect")
	dbname := os.Getenv("dbname")
	host := os.Getenv("host")
	user := os.Getenv("user")
	password := os.Getenv("password")

	dbURL := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode = disable", host, user, dbname, password)

	db, err := gorm.Open(dialect, dbURL)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}

func MigrateModel(db *gorm.DB) {
	db.Debug().AutoMigrate(&entity.Customer{})
	db.Debug().AutoMigrate(&entity.Catagory{})
	db.Debug().AutoMigrate(&entity.Item{})
	db.Debug().AutoMigrate(&entity.Record{})
	db.Debug().AutoMigrate(&entity.CartInfo{})
	db.Debug().AutoMigrate(&entity.ItemInfo{})
}
