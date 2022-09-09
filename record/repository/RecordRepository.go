package recordRepository

import (
	//"fmt"

	//"fmt"
	//"log"

	"github.com/jinzhu/gorm"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	"github.com/motikingo/ecommerceRESTAPI-Go/record"
)

type RecordRepo struct {
	db *gorm.DB
}

func NewRecordRepo(db *gorm.DB) record.RecordRepository {
	return &RecordRepo{db: db}
}

func (recoRepo *RecordRepo) GetRecords() []entity.Record {

	var records []entity.Record

	if err := recoRepo.db.Find(&records).GetErrors(); len(err) > 0 {
		return nil
	}
	return records

}

// func(recoRepo *RecordRepo)GetRecord(id uint)(*entity.Record,[]error){

// 	var record entity.Record
// 	err := recoRepo.db.First(&record).GetErrors()
// 	if len(err)>0 {
// 		return nil,err
// 	}
// 	return &record,err
// }
func (recoRepo *RecordRepo) GetRecordByUserID(user_id uint) *entity.Record {
	var record entity.Record

	if err := recoRepo.db.Preload("Cart_Infos.Item_Infos").First(&record, user_id).GetErrors(); len(err) > 0 {
		//log.Println(err)
		return nil
	}

	return &record
}

func (recoRepo *RecordRepo) CreateInfo(info entity.CartInfo) *entity.CartInfo {

	if ers := recoRepo.db.Create(&info).GetErrors(); len(ers) > 0 {
		return nil
	}
	return &info
}

func (recoRepo *RecordRepo) UpdateRecord(reco entity.Record) *entity.Record {

	record := recoRepo.GetRecordByUserID(reco.UserId)

	if record == nil {
		return nil
	}

	for _, itmI := range reco.Cart_Infos[len(reco.Cart_Infos)-1].Item_Infos {
		if ers := recoRepo.db.Create(&itmI).GetErrors(); len(ers) > 0 {
			return nil
		}
	}
	var carts entity.CartInfo

	recoRepo.db.First(&carts, reco.Cart_Infos[len(reco.Cart_Infos)-1].ID)

	recoRepo.db.Model(&record).Association("Item_Infos").Clear()

	carts.Item_Infos = reco.Cart_Infos[len(reco.Cart_Infos)-1].Item_Infos

	if ers := recoRepo.db.Save(&carts).GetErrors(); len(ers) > 0 {

		return nil
	}
	recoRepo.db.Model(&record).Association("Cart_Infos").Clear()
	record.Cart_Infos = reco.Cart_Infos
	//fmt.Println(reco.Cart_Infos)
	if ers := recoRepo.db.Save(&record).GetErrors(); len(ers) > 0 {
		return nil
	}
	return record
}
func (recoRepo *RecordRepo) CreateRecord(reco entity.Record) *entity.Record {
	//record := reco

	if ers := recoRepo.db.Create(&reco).GetErrors(); len(ers) > 0 {
		return nil
	}
	return &reco
}
func (recoRepo *RecordRepo) ClearRecord(id uint) *entity.Record {
	//record,ers:= recoRepo.GetRecord(id)
	var record entity.Record
	if ers := recoRepo.db.Delete(&record, id).GetErrors(); len(ers) > 0 {
		return nil
	}
	return &record
}
