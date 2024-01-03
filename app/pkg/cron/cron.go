package cron

import (
	"github.com/robfig/cron"
	"github.com/uussoop/idp-go/pkg/cron/jobs"
)

var CronJob *cron.Cron

func init() {
	CronJob = cron.New()

	initJobs()
}

func initJobs() {
	if CronJob != nil {
		// every 7 days get models
		CronJob.AddFunc("0 0 0 ? * 7/7 *", jobs.RefreshKeys)

	}
}

func Start() {
	if CronJob != nil {
		CronJob.Start()
	}
}
