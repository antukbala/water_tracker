package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {
	beego.GlobalControllerRouter["water_tracker/controllers:WaterController"] = append(beego.GlobalControllerRouter["water_tracker/controllers:WaterController"],
		beego.ControllerComments{
			Method:           "TotalWaterConsumedByIDLastNDays",
			Router:           "/total_water_consumed_last_ndays",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["water_tracker/controllers:WaterController"] = append(beego.GlobalControllerRouter["water_tracker/controllers:WaterController"],
		beego.ControllerComments{
			Method:           "TotalWaterConsumedByID",
			Router:           "/total_water_consumed_today",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["water_tracker/controllers:WaterController"] = append(beego.GlobalControllerRouter["water_tracker/controllers:WaterController"],
		beego.ControllerComments{
			Method:           "ShowAllWaterDetails",
			Router:           "/show_all_water_details",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["water_tracker/controllers:WaterController"] = append(beego.GlobalControllerRouter["water_tracker/controllers:WaterController"],
		beego.ControllerComments{
			Method:           "ShowWaterDetails",
			Router:           "/show_water_details",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["water_tracker/controllers:WaterController"] = append(beego.GlobalControllerRouter["water_tracker/controllers:WaterController"],
		beego.ControllerComments{
			Method:           "DeleteWaterInput",
			Router:           "/delete_water_input",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["water_tracker/controllers:WaterController"] = append(beego.GlobalControllerRouter["water_tracker/controllers:WaterController"],
		beego.ControllerComments{
			Method:           "EditWaterInput",
			Router:           "/edit_water_input",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["water_tracker/controllers:WaterController"] = append(beego.GlobalControllerRouter["water_tracker/controllers:WaterController"],
		beego.ControllerComments{
			Method:           "InputWater",
			Router:           "/input_water",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["water_tracker/controllers:ObjectController"] = append(beego.GlobalControllerRouter["water_tracker/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["water_tracker/controllers:ObjectController"] = append(beego.GlobalControllerRouter["water_tracker/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["water_tracker/controllers:ObjectController"] = append(beego.GlobalControllerRouter["water_tracker/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:objectId",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["water_tracker/controllers:ObjectController"] = append(beego.GlobalControllerRouter["water_tracker/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:objectId",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["water_tracker/controllers:ObjectController"] = append(beego.GlobalControllerRouter["water_tracker/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:objectId",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["water_tracker/controllers:UserController"] = append(beego.GlobalControllerRouter["water_tracker/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["water_tracker/controllers:UserController"] = append(beego.GlobalControllerRouter["water_tracker/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["water_tracker/controllers:UserController"] = append(beego.GlobalControllerRouter["water_tracker/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["water_tracker/controllers:UserController"] = append(beego.GlobalControllerRouter["water_tracker/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["water_tracker/controllers:UserController"] = append(beego.GlobalControllerRouter["water_tracker/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["water_tracker/controllers:UserController"] = append(beego.GlobalControllerRouter["water_tracker/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           "/login",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["water_tracker/controllers:UserController"] = append(beego.GlobalControllerRouter["water_tracker/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           "/logout",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
