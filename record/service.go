package record

import "github.com/motikingo/ecommerceRESTAPI-Go/entity"

type RecordService interface{
	GetRecords()([]entity.Record)
	//GetRecord(id uint)(*entity.Record)
	GetRecordByUserID(user_Id uint)*entity.Record
	UpdateRecord(car entity.Record)(*entity.Record)
	CreateRecord(reco entity.Record)(*entity.Record)
	CreateInfo(info entity.CartInfo)(*entity.CartInfo)
	ClearRecord(id uint)(*entity.Record)

}
