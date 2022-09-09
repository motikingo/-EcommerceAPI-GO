package itemRepository

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	"github.com/motikingo/ecommerceRESTAPI-Go/item"
)

type ItemRepo struct {
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) item.ItemRepository {
	return &ItemRepo{db: db}
}

func (itemRepo *ItemRepo) GetItems() ([]entity.Item, []error) {
	var items []entity.Item
	ers := itemRepo.db.Preload("Catagories").Find(&items).GetErrors()
	if len(ers) > 0 {
		return nil, ers
	}
	return items, nil
}
func (itemRepo *ItemRepo) GetItem(id uint) (*entity.Item, []error) {

	var item entity.Item
	ers := itemRepo.db.Preload("Catagories").First(&item, id).GetErrors()
	if len(ers) > 0 {
		return nil, ers
	}
	return &item, nil
}

func (itemRepo *ItemRepo) IsItemNameExist(name string) *entity.Item {

	var item entity.Item
	ers := itemRepo.db.First(&item, name).GetErrors()
	if len(ers) > 0 {
		return nil
	}
	return &item
}
func (itemRepo *ItemRepo) UpdateItem(item entity.Item) (*entity.Item, []error) {

	itm, e := itemRepo.GetItem(item.ID)

	if len(e) > 0 {
		return nil, e
	}
	itm.Name = func() string {
		if itm.Name != item.Name {
			return item.Name
		}
		return itm.Name
	}()
	itm.Description = func() string {
		if item.Description != itm.Description {
			return item.Description
		}
		return itm.Description
	}()

	itm.Brand = func() string {
		if item.Brand != itm.Brand {
			return item.Brand
		}
		return itm.Brand
	}()
	itm.Price = func() float64 {
		if item.Price != item.Price {
			return item.Price
		}
		return item.Price
	}()
	itm.Image = func() string {
		if item.Image != itm.Image {
			return item.Image
		}
		return itm.Image
	}()
	itm.Number = func() int {
		if item.Number != itm.Number {
			return item.Number
		}
		return itm.Number
	}()
	itm.ProductionDate = func() time.Time {
		if !item.ProductionDate.Equal(itm.ProductionDate) {
			return item.ProductionDate
		}
		return itm.ProductionDate
	}()
	itm.ExpireDate = func() time.Time {
		if !item.ExpireDate.Equal(item.ExpireDate) {
			return item.ExpireDate
		}
		return itm.ExpireDate
	}()
	itm.Catagories = func() []entity.Catagory {
		if !reflect.DeepEqual(itm.Catagories, item.Catagories) {
			itemRepo.db.Model(itm).Association("Catagories").Clear()
			return item.Catagories
		}
		return itm.Catagories
	}()

	if ers := itemRepo.db.Save(&itm).GetErrors(); len(ers) > 0 {
		return nil, ers
	}
	return &item, nil
}
func (itemRepo *ItemRepo) CreateItem(item entity.Item) (*entity.Item, []error) {
	itm, ers := itemRepo.GetItem(item.ID)
	fmt.Println("Yooo")
	if itm != nil && len(ers) == 0 {
		log.Fatal("Item with this Id already exist")
		return nil, nil
	} else if itm == nil && len(ers) > 0 {
		fmt.Println()
		if ers := itemRepo.db.Create(&item).GetErrors(); len(ers) > 0 {
			return nil, ers
		}
		return &item, nil

	}
	return nil, nil

}
func (itemRepo *ItemRepo) DeleteItem(id uint) (*entity.Item, []error) {

	item, ers := itemRepo.GetItem(id)
	if item == nil && len(ers) > 0 {
		log.Fatal("this user doesn't exist")
		return nil, ers
	} else if item != nil && len(ers) == 0 {

		if ers := itemRepo.db.Delete(item, id).GetErrors(); len(ers) > 0 {
			return nil, ers
		}
		itemRepo.db.Model(item).Association("Catagories").Clear()
		return item, nil
	}

	return nil, nil

}
