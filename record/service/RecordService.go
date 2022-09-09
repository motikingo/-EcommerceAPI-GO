package recordService

import (
	"github.com/motikingo/ecommerceRESTAPI-Go/entity"
	"github.com/motikingo/ecommerceRESTAPI-Go/record"
)

type RecordServ struct {
	repo record.RecordRepository
}

func NewRecordServ(repo record.RecordRepository) record.RecordService {
	return &RecordServ{repo: repo}
}

func (carServ *RecordServ) GetRecords() []entity.Record {

	return carServ.repo.GetRecords()

}

// func(carServ *RecordServ)GetRecord(id uint)(*entity.Record,[]error){

// 	//var cart entity.Cart
// 	record,err := carServ.repo.GetRecord(id)
// 	if len(err)>0 {
// 		return nil,err
// 	}
// 	return record,err
// }

func (carServ *RecordServ) GetRecordByUserID(user_Id uint) *entity.Record {
	return carServ.repo.GetRecordByUserID(user_Id)
}

func (carServ *RecordServ) UpdateRecord(reco entity.Record) *entity.Record {

	return carServ.repo.UpdateRecord(reco)
}
func (carServ *RecordServ) CreateRecord(reco entity.Record) *entity.Record {

	return carServ.repo.CreateRecord(reco)
}
func (carServ *RecordServ) CreateInfo(info entity.CartInfo) *entity.CartInfo {
	return carServ.repo.CreateInfo(info)
}

func (carServ *RecordServ) ClearRecord(id uint) *entity.Record {

	return carServ.repo.ClearRecord(id)
}
