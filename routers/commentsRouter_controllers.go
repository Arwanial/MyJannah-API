package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:CategoryFacilityController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:CategoryFacilityController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:CategoryFacilityController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:CategoryFacilityController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:CategoryFacilityController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:CategoryFacilityController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:CategoryFacilityController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:CategoryFacilityController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:CategoryFacilityController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:CategoryFacilityController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:FasilitasController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:FasilitasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:FasilitasController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:FasilitasController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:FasilitasController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:FasilitasController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:FasilitasController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:FasilitasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:FasilitasController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:FasilitasController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:JamaahController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:JamaahController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:JamaahController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:JamaahController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:JamaahController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:JamaahController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:JamaahController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:JamaahController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:JamaahController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:JamaahController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ManasikScheduleController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ManasikScheduleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ManasikScheduleController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ManasikScheduleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ManasikScheduleController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ManasikScheduleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ManasikScheduleController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ManasikScheduleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ManasikScheduleController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ManasikScheduleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:MenuController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:MenuController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:MenuController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:MenuController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:MenuController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:MenuController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:MenuController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:MenuController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:MenuController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:MenuController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:PaketTravelController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:PaketTravelController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:PaketTravelController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:PaketTravelController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:PaketTravelController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:PaketTravelController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:PaketTravelController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:PaketTravelController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:PaketTravelController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:PaketTravelController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:RegistrationController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:RegistrationController"],
		beego.ControllerComments{
			Method: "RegisterTravel",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:RoleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:RoleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ScheduleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ScheduleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ScheduleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ScheduleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:ScheduleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:SubPaketTravelController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:SubPaketTravelController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:SubPaketTravelController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:SubPaketTravelController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:SubPaketTravelController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:SubPaketTravelController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:SubPaketTravelController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:SubPaketTravelController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:SubPaketTravelController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:SubPaketTravelController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:TravelagentController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:TravelagentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:TravelagentController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:TravelagentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:TravelagentController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:TravelagentController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:TravelagentController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:TravelagentController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:TravelagentController"] = append(beego.GlobalControllerRouter["github.com/Arwanial/MyJannah-API/controllers:TravelagentController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
