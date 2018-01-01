// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/Arwanial/MyJannah-API/controllers"

	"github.com/astaxie/beego"
)

func init() {
beego.Router("/v1/travelagent/register", &controllers.RegistrationController{}, "post:RegisterTravel")
beego.Router("/v1/travelagent/login", &controllers.LoginController{}, "post:Login")
beego.Router("/v1/travelagent/resendregister", &controllers.ResendRegisterNotifController{}, "post:ResendNotifReg")
}
