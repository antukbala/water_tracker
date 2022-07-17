package models

var query = make(map[string]string)

func init() {
	query["insertWaterAmount"] = "insert into personal.water_tracker(user_id,water_amount,note,input_time) values(?,?,?,now());"
	query["checkProcessIdExistance"] = "select process_id from personal.water_tracker where process_id=?;"
	query["updateWaterInputByProcessId"] = "update personal.water_tracker set user_id=?, water_amount=?, note=?, update_time=now() where process_id=?;"
	query["getWaterDetailsByProcessId"] = "select process_id, user_id, water_amount, note from personal.water_tracker where process_id=?;"
	query["deleteWaterDetailsByProcessId"] = "delete from personal.water_tracker where process_id=?;"
	query["getAllWaterDetails"] = "select process_id, user_id, water_amount, note from personal.water_tracker order by process_id desc;"
	query["totalWaterConsumedTodayByUserId"] = "select sum(water_amount) as total_water from personal.water_tracker where user_id=? and date(input_time)=date(now());"
	query["totalWaterConsumedByUserIdLastNdays"] = "select sum(water_amount) as total_water from personal.water_tracker where user_id=? and date(input_time)>=CURDATE()-INTERVAL ? DAY;"
}

func SqlQuery(key string) string {
	return query[key]
}
