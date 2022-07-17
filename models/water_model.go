package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type InsertWaterAmountBody struct {
	UserId      int64   `json:"user_id"`
	WaterAmount float64 `json:"water_amount"`
	Note        string  `json:"note"`
}

type InsertWaterAmountResponse struct {
	Message   string `json:"message"`
	ProcessId int64  `json:"process_id"`
}

type EditWaterAmountResponse struct {
	Message     string  `json:"message"`
	ProcessId   int64   `json:"process_id"`
	UserId      int64   `json:"user_id"`
	WaterAmount float64 `json:"water_amount"`
	Note        string  `json:"note"`
}

type GetAllWaterResponse struct {
	ProcessId   int64   `json:"process_id"`
	UserId      int64   `json:"user_id"`
	WaterAmount float64 `json:"water_amount"`
	Note        string  `json:"note"`
}

type GetAllWaterControllerResponse struct {
	TotalInstanse int64                 `json:"total_items"`
	Output        []GetAllWaterResponse `json:"output"`
}

type WaterTrackerDetails struct {
	ProcessId   int64     `json:"process_id" orm:"pk; column(process_id)"`
	UserId      int64     `json:"user_id" orm:"column(user_id)"`
	WaterAmount float64   `json:"water_amount" orm:"column(water_amount)"`
	Note        string    `json:"note" orm:"column(note)"`
	InputTime   time.Time `json:"input_time" orm:"type(datetime)"`
	UpdateTime  time.Time `json:"update_time" orm:"type(datetime)"`
}

type WaterTracker struct {
	ProcessId   int64     `json:"process_id" orm:"pk; column(process_id)"`
	UserId      int64     `json:"user_id" orm:"column(user_id)"`
	WaterAmount float64   `json:"water_amount" orm:"column(water_amount)"`
	Note        string    `json:"note" orm:"column(note)"`
	InputTime   time.Time `json:"-" orm:"type(datetime)"`
	UpdateTime  time.Time `json:"-" orm:"type(datetime)"`
}

type InputWaterPostPayload struct {
	UserId      int64   `json:"user_id" valid:"Required"`
	WaterAmount float64 `json:"water_amount" valid:"Required"`
	Note        string  `json:"note" valid:"Required"`
}

type DeleteWaterInputPostPayload struct {
	ProcessId int64 `json:"process_id" valid:"Required"`
}

type UserIdPayload struct {
	UserId int64 `json:"user_id" valid:"Required"`
}

type UserIdAndDaysPayload struct {
	UserId int64 `json:"user_id" valid:"Required"`
	Days   int64 `json:"days" valid:"Required"`
}

type TotalWaterConsumedByIDLastNDays struct {
	UserId int64 `json:"user_id" valid:"Required"`
	Days   int64 `json:"days" valid:"Required"`
}

type TotalWaterConsumedByIDResponse struct {
	TotalWater float64 `json:"total_water"`
}

type EditWaterInputPostPayload struct {
	ProcessId   int64   `json:"process_id" valid:"Required"`
	UserId      int64   `json:"user_id" valid:"Required"`
	WaterAmount float64 `json:"water_amount" valid:"Required"`
	Note        string  `json:"note" valid:"Required"`
}

type OutputTotalWaterConsumedByIdToday struct {
	TotalWater float64 `json:"total_water"`
	Date       string  `json:"date"`
}

type OutputTotalWaterConsumedByIdLastNdays struct {
	UserId         int64   `json:"user_id"`
	Days           int64   `json:"days"`
	TotalWater     float64 `json:"total_water"`
	AvgWaterPerDay float64 `json:"avg_water_per_day"`
}

func init() {
	orm.RegisterModel(new(WaterTracker))
}

// type WithdrawRequest struct {
// 	WithdrawRequestID int       `json:"withdraw_request_id" orm:"pk; column(withdraw_request_id)"`
// 	LedgerId          int       `json:"ledger_id" orm:"column(ledger_id)"`
// 	LedgerPhone       string    `json:"ledger_phone" orm:"column(ledger_phone)"`
// 	PleaderId         int       `json:"pleader_id" orm:"column(pleader_id)"`
// 	PleaderType       string    `json:"pleader_type" orm:"column(pleader_type)"`
// 	IsActive          string    `json:"is_active" orm:"column(is_active)"`
// 	ReceiverId        int       `json:"receiver_id" orm:"column(receiver_id)"`
// 	ReceiverName      string    `json:"receiver_name" orm:"column(receiver_name)"`
// 	ReceiverPhone     string    `json:"receiver_phone" orm:"column(receiver_phone)"`
// 	ReceiverType      string    `json:"receiver_type" orm:"column(receiver_type)"`
// 	RequestedAmount   int       `json:"requested_amount" orm:"column(requested_amount)"`
// 	CreatedBy         string    `json:"created_by" orm:"column(created_by)"`
// 	UpdatedBy         string    `json:"updated_by" orm:"column(updated_by)"`
// 	CreateDate        time.Time `json:"-" orm:"auto_now_add;type(datetime)"`
// 	UpdateDate        time.Time `json:"-" orm:"auto_now_add;type(datetime)"`
// 	WithdrawMethod    string    `json:"withdraw_method" orm:"column(withdraw_method)"`
// }

// func InsertWaterAmount(body WaterTracker) (response WaterTrackerDetails, err error) {
// 	o := orm.NewOrm()
// 	processId, err := o.Insert(&body)
// 	if err != nil {
// 		return response, err
// 	}
// 	response.ProcessId = int(processId)
// 	return response, nil
// }

func InsertWaterAmount(values InputWaterPostPayload) (response InsertWaterAmountResponse, err error) {
	o := orm.NewOrm()
	res, err := o.Raw(SqlQuery("insertWaterAmount"), values.UserId, values.WaterAmount, values.Note).Exec()

	if err != nil {
		return response, err
	}
	response.ProcessId, _ = res.LastInsertId()
	return response, nil
}

func CheckProcessIdExistance(processId int64) int64 {
	o := orm.NewOrm()
	var value []orm.Params
	response, _ := o.Raw(SqlQuery("checkProcessIdExistance"), processId).Values(&value)
	return response
}

func UpdateWaterInputByProcessId(values EditWaterInputPostPayload) int64 {
	o := orm.NewOrm()
	response, _ := o.Raw(SqlQuery("updateWaterInputByProcessId"), values.UserId, values.WaterAmount, values.Note, values.ProcessId).Exec()
	affectedRows, _ := response.RowsAffected()
	return affectedRows
}

func GetWaterDetailsByProcessId(processId int64) (map[string]interface{}, error) {
	o := orm.NewOrm()
	var value []orm.Params
	if _, err := o.Raw(SqlQuery("getWaterDetailsByProcessId"), processId).Values(&value); err != nil {
		return nil, err
	}
	return value[0], nil
}

func GetAllWaterDetails() ([]GetAllWaterResponse, error) {
	o := orm.NewOrm()
	var response []GetAllWaterResponse
	if _, err := o.Raw(SqlQuery("getAllWaterDetails")).QueryRows(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func GetTotalWaterConsumedByID(values UserIdPayload) ([]TotalWaterConsumedByIDResponse, error) {
	o := orm.NewOrm()
	var response []TotalWaterConsumedByIDResponse
	if _, err := o.Raw(SqlQuery("totalWaterConsumedTodayByUserId"), values.UserId).QueryRows(&response); err != nil {
		return response, err
	}
	return response, nil
}

func GetTotalWaterConsumedByIDLastNdays(values UserIdAndDaysPayload) ([]TotalWaterConsumedByIDResponse, error) {
	o := orm.NewOrm()
	var response []TotalWaterConsumedByIDResponse
	if _, err := o.Raw(SqlQuery("totalWaterConsumedByUserIdLastNdays"), values.UserId, values.Days).QueryRows(&response); err != nil {
		return response, err
	}
	return response, nil
}

func DeleteWaterDetailsByProcessId(processId int64) int64 {
	o := orm.NewOrm()
	response, _ := o.Raw(SqlQuery("deleteWaterDetailsByProcessId"), processId).Exec()
	affectedRows, _ := response.RowsAffected()
	return affectedRows
}
