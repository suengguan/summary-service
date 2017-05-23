package service

import (
	"fmt"
	"model"

	appApi "api/app_service"
	daoApi "api/dao_service"

	"github.com/astaxie/beego"

	"encoding/json"
)

type SummaryService struct {
}

func (this *SummaryService) Get(userId int64) (*model.Summary, error) {
	var err error
	var summary model.Summary

	// get all projects
	beego.Debug("->get all projects")
	var projects []*model.Project
	projects, err = daoApi.BussinessDaoApi.GetAllProjects(userId)
	if err != nil {
		beego.Debug("get project by user failed")
		err = fmt.Errorf("%s", "get all project failed")
		return nil, err
	}
	summary.ProjectTotalCount = len(projects)

	// get all jobs status
	var jobCnt int
	beego.Debug("->get all jobs status")
	if len(projects) > 0 {
		for _, p := range projects {
			jobCnt += len(p.Jobs)
		}
	}
	summary.JobTotalCount = jobCnt

	if jobCnt > 0 {
		var allJobStatus []*model.JobStatus
		var result []byte
		result, err = appApi.StatusApi.GetAllJobStatus(userId)
		if err != nil {
			beego.Debug("get all job status failed!")
			return nil, err
		}
		//beego.Debug(string(result))

		err = json.Unmarshal(result, &allJobStatus)
		if err != nil {
			beego.Debug("Unmarshal data failed")
			return nil, err
		}

		for _, status := range allJobStatus {
			if status.Status == model.JOB_STATUS_TYPE_RUN || status.Status == model.JOB_STATUS_TYPE_READY {
				summary.JobRunningCount += 1
			} else if status.Status == model.JOB_STATUS_TYPE_FINISH {
				summary.JobSuccessCount += 1
			} else if status.Status == model.JOB_STATUS_TYPE_ERROR {
				summary.JobFailedCount += 1
			}
		}
	}

	// get resource
	beego.Debug("->get resource")
	var resource *model.Resource
	resource, err = daoApi.ResourceDaoApi.GetByUserId(userId)
	if err != nil {
		beego.Debug("get resource failed")
		return nil, err
	}

	summary.CpuTotal = resource.CpuTotalResource
	summary.CpuUsed = resource.CpuUsageResource
	summary.CpuUsedPercent = (int)(resource.CpuUsageResource / resource.CpuTotalResource * 100.0)
	summary.CpuUnit = resource.CpuUnit
	summary.MemoryTotal = resource.MemoryTotalResource
	summary.MemoryUsed = resource.MemoryUsageResource
	summary.MemoryUsedPercent = (int)(resource.MemoryUsageResource / resource.MemoryTotalResource * 100.0)
	summary.MemoryUnit = resource.MemoryUnit

	beego.Debug("result:", summary)

	return &summary, err
}
