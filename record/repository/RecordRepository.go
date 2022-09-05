package recordRepository
import (
	"github.com/jinzhu/gorm"
	"github.com/motikingo/ecommerceRESTAPI-Go/record"
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
)

type RecordRepo struct{
	db *gorm.DB
}

func NewRecordRepo(db *gorm.DB) record.RecordRepository{
	return &RecordRepo{db: db}
}

func(recoRepo *RecordRepo) GetRecords()([]entity.Record){

	var records []entity.Record
	
	if err := recoRepo.db.Find(&records).GetErrors(); len(err)>0 {
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
func(recoRepo *RecordRepo)GetRecordByUserID(user_id uint)*entity.Record{
	var record entity.Record
	err := recoRepo.db.Preload("Carts").First(&record,user_id).GetErrors()
	if len(err)>0{
		return nil
	}
	return &record
}

func(carRepo *RecordRepo)UpdateRecord(reco entity.Record)(*entity.Record){

	record:= carRepo.GetRecordByUserID(reco.UserId)

	if record == nil {
		return nil
	}
	record.Carts = reco.Carts
	
	if ers := carRepo.db.Save(&record).GetErrors(); len(ers)>0 {
		return nil
	}
	return record
}
func(recoRepo *RecordRepo)CreateRecord(reco entity.Record)(*entity.Record){
	//record := reco
	if ers := recoRepo.db.Create(&reco).GetErrors();len(ers)>0 {
		return nil
	}
	return &reco
}
func(recoRepo *RecordRepo)ClearRecord(id uint)(*entity.Record){
	//record,ers:= recoRepo.GetRecord(id)
	var record entity.Record
	if ers := recoRepo.db.Delete(&record,id).GetErrors();len(ers)>0 {
		return nil
	}
	return &record
}



