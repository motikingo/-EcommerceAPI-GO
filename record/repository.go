package record

import "github.com/motikingo/ecommerceRESTAPI-Go/entity"

type RecordRepository interface{
	GetRecords()([]entity.Record)
	//GetRecord(id uint)(*entity.Record)
	GetRecordByUserID(user_Id uint)*entity.Record
	UpdateRecord(car entity.Record)(*entity.Record)
	CreateRecord(reco entity.Record)(*entity.Record)
	ClearRecord(id uint)(*entity.Record)
}
