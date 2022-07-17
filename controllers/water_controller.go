package controllers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
	"water_tracker/models"
	services "water_tracker/services"

	beego "github.com/beego/beego/v2/server/web"
)

func ProcessIdChecking(processId int64, w *WaterController) int64 {
	if process := models.CheckProcessIdExistance(processId); process == 0 {
		w.Data["json"] = services.APIResponse{Code: 1005, Message: "Process Id not found"}
		w.ServeJSON()
		return process
	}
	return 1
}

// Operations about Users
type WaterController struct {
	beego.Controller
}

// @router /input_water [post]
func (w *WaterController) InputWater() {
	payload := models.InputWaterPostPayload{}
	json.Unmarshal(w.Ctx.Input.RequestBody, &payload)

	err := services.BodyParamCheck(&payload)
	if err != nil {
		w.Data["json"] = services.APIResponse{Code: 2000, Message: err.Error()}
		w.ServeJSON()
		return
	}

	// allData := models.WaterTracker{
	// 	UserId:      payload.UserId,
	// 	WaterAmount: payload.WaterAmount,
	// 	Note:        payload.Note,
	// }

	response, err := models.InsertWaterAmount(payload)
	if err != nil {
		response.Message = "Error occured while inserting data"
		w.Data["json"] = response
		w.ServeJSON()
		return
	}

	response.Message = "Successafully inserted data"
	w.Data["json"] = response
	w.ServeJSON()
}

// @router /edit_water_input [post]
func (w *WaterController) EditWaterInput() {
	payload := models.EditWaterInputPostPayload{}
	json.Unmarshal(w.Ctx.Input.RequestBody, &payload)

	err := services.BodyParamCheck(&payload)
	if err != nil {
		w.Data["json"] = services.APIResponse{Code: 2000, Message: err.Error()}
		w.ServeJSON()
		return
	}

	if check := ProcessIdChecking(payload.ProcessId, w); check == 0 {
		return
	}

	if affectedRows := models.UpdateWaterInputByProcessId(payload); affectedRows == 0 {
		w.Data["json"] = services.APIResponse{Code: 1006, Message: "Error Occured when updating"}
		w.ServeJSON()
		return
	}

	x, err := models.GetWaterDetailsByProcessId(payload.ProcessId)
	if err != nil {
		w.Data["json"] = services.APIResponse{Code: 2110, Message: err.Error()}
		w.ServeJSON()
		return
	}

	var output models.EditWaterAmountResponse
	output.Message = "Successfully updated water input"
	output.ProcessId, _ = strconv.ParseInt(x["process_id"].(string), 10, 8)
	output.UserId, _ = strconv.ParseInt(x["user_id"].(string), 10, 8)
	output.WaterAmount, _ = strconv.ParseFloat(x["water_amount"].(string), 8)
	output.Note = x["note"].(string)

	w.Data["json"] = output
	w.ServeJSON()
}

// @router /delete_water_input [post]
func (w *WaterController) DeleteWaterInput() {
	payload := models.DeleteWaterInputPostPayload{}
	json.Unmarshal(w.Ctx.Input.RequestBody, &payload)

	err := services.BodyParamCheck(&payload)
	if err != nil {
		w.Data["json"] = services.APIResponse{Code: 2000, Message: err.Error()}
		w.ServeJSON()
		return
	}

	if check := ProcessIdChecking(payload.ProcessId, w); check == 0 {
		return
	}

	if affertedRows := models.DeleteWaterDetailsByProcessId(payload.ProcessId); affertedRows == int64(1) {
		var output models.InsertWaterAmountResponse
		output.Message = "Successfully deleted water input"
		output.ProcessId = payload.ProcessId
		w.Data["json"] = output
	} else {
		w.Data["json"] = services.APIResponse{Code: 1006, Message: "Error Occured when deleting"}
	}
	w.ServeJSON()
}

// @router /show_water_details [post]
func (w *WaterController) ShowWaterDetails() {
	payload := models.DeleteWaterInputPostPayload{}
	json.Unmarshal(w.Ctx.Input.RequestBody, &payload)

	err := services.BodyParamCheck(&payload)
	if err != nil {
		w.Data["json"] = services.APIResponse{Code: 2000, Message: err.Error()}
		w.ServeJSON()
		return
	}

	if check := ProcessIdChecking(payload.ProcessId, w); check == 0 {
		return
	}

	x, err := models.GetWaterDetailsByProcessId(payload.ProcessId)
	if err != nil {
		w.Data["json"] = services.APIResponse{Code: 2110, Message: err.Error()}
		w.ServeJSON()
		return
	}

	var output models.EditWaterAmountResponse
	output.Message = "Successfully fetched water input details"
	output.ProcessId, _ = strconv.ParseInt(x["process_id"].(string), 10, 8)
	output.UserId, _ = strconv.ParseInt(x["user_id"].(string), 10, 8)
	output.WaterAmount, _ = strconv.ParseFloat(x["water_amount"].(string), 8)
	output.Note = x["note"].(string)

	w.Data["json"] = output
	w.ServeJSON()
}

// @router /show_all_water_details [post]
func (w *WaterController) ShowAllWaterDetails() {
	response, err := models.GetAllWaterDetails()
	if err != nil {
		w.Data["json"] = services.APIResponse{Code: 2110, Message: err.Error()}
		w.ServeJSON()
		return
	}
	output := models.GetAllWaterControllerResponse{
		TotalInstanse: int64(len(response)),
		Output:        response,
	}
	w.Data["json"] = output
	w.ServeJSON()
}

// @router /total_water_consumed_today [post]
func (w *WaterController) TotalWaterConsumedByID() {
	payload := models.UserIdPayload{}
	json.Unmarshal(w.Ctx.Input.RequestBody, &payload)

	err := services.BodyParamCheck(&payload)
	if err != nil {
		w.Data["json"] = services.APIResponse{Code: 2000, Message: err.Error()}
		w.ServeJSON()
		return
	}

	response, err := models.GetTotalWaterConsumedByID(payload)
	if err != nil {
		w.Data["json"] = services.APIResponse{Code: 2110, Message: err.Error()}
		w.ServeJSON()
		return
	}
	currentTime := time.Now()
	year, month, date := currentTime.Date()
	today := fmt.Sprintf("%v-%v-%v", date, month, year)
	totalWater := response[0]
	fmt.Println(totalWater.TotalWater, reflect.TypeOf(totalWater.TotalWater))
	w.Data["json"] = models.OutputTotalWaterConsumedByIdToday{TotalWater: response[0].TotalWater, Date: today}
	w.ServeJSON()
}

// @router /total_water_consumed_last_ndays [post]
func (w *WaterController) TotalWaterConsumedByIDLastNDays() {
	payload := models.UserIdAndDaysPayload{}
	json.Unmarshal(w.Ctx.Input.RequestBody, &payload)

	err := services.BodyParamCheck(&payload)
	if err != nil {
		w.Data["json"] = services.APIResponse{Code: 2000, Message: err.Error()}
		w.ServeJSON()
		return
	}

	response, err := models.GetTotalWaterConsumedByIDLastNdays(payload)
	if err != nil {
		w.Data["json"] = services.APIResponse{Code: 2110, Message: err.Error()}
		w.ServeJSON()
		return
	}

	fmt.Println(response)
	totalWater := response[0].TotalWater
	avgWater := totalWater / float64(payload.Days)
	w.Data["json"] = models.OutputTotalWaterConsumedByIdLastNdays{UserId: payload.UserId, Days: payload.Days, TotalWater: totalWater, AvgWaterPerDay: avgWater}
	w.ServeJSON()
}

// @router /total_water_consumed_from_date [post]
func (w *WaterController) TotalWaterConsumedByIDMultipleDays() {
	payload := models.UserIdPayload{}
	json.Unmarshal(w.Ctx.Input.RequestBody, &payload)

	err := services.BodyParamCheck(&payload)
	if err != nil {
		w.Data["json"] = services.APIResponse{Code: 2000, Message: err.Error()}
		w.ServeJSON()
		return
	}

	response, err := models.GetTotalWaterConsumedByID(payload)
	if err != nil {
		w.Data["json"] = services.APIResponse{Code: 2110, Message: err.Error()}
		w.ServeJSON()
		return
	}
	currentTime := time.Now()
	year, month, date := currentTime.Date()
	today := fmt.Sprintf("%v-%v-%v", date, month, year)
	totalWater := response[0]
	fmt.Println(totalWater.TotalWater, reflect.TypeOf(totalWater.TotalWater))
	w.Data["json"] = models.OutputTotalWaterConsumedByIdToday{TotalWater: response[0].TotalWater, Date: today}
	w.ServeJSON()
}
