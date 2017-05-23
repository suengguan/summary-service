package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["app-service/summary-service/controllers:SummaryController"] = append(beego.GlobalControllerRouter["app-service/summary-service/controllers:SummaryController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:userId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
