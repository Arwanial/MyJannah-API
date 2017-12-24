// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"MyJannah-API/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/menu",
			beego.NSInclude(
				&controllers.MenuController{},
			),
		),

		beego.NSNamespace("/role",
			beego.NSInclude(
				&controllers.RoleController{},
			),
		),

		beego.NSNamespace("/travelagent",
			beego.NSInclude(
				&controllers.TravelagentController{},
			),
		),

		beego.NSNamespace("/registertravelagent",
			beego.NSInclude(
				&controllers.RegistrationController{},
			),
		),

		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
