package controllers

import (
	"app-service/summary-service/models"
	"app-service/summary-service/service"
	"encoding/json"
	"fmt"
	"model"

	"github.com/astaxie/beego"
)

// Operations about Summary
type SummaryController struct {
	beego.Controller
}

// @Title Get
// @Description get summary
// @Param	userid		path 	int64	true		"The key for staticblock"
// @Success 200 {object} models.Response
// @router /:userId [get]
func (this *SummaryController) Get() {
	var err error
	var response models.Response

	var userId int64
	userId, err = this.GetInt64(":userId")
	// beego.Debug("Get", userId)
	if userId > 0 && err == nil {
		var svc service.SummaryService
		var summary *model.Summary
		var result []byte
		summary, err = svc.Get(userId)
		if err == nil {
			result, err = json.Marshal(summary)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug(err)
		err = fmt.Errorf("%s", "user id is invalid")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}
	this.Data["json"] = &response

	this.ServeJSON()
}
