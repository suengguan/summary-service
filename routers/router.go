// @APIVersion 1.0.0
// @Title summary-service API
// @Description summary-service only serve summary
// @Contact qsg@corex-tek.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"app-service/summary-service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/summary",
			beego.NSInclude(
				&controllers.SummaryController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
