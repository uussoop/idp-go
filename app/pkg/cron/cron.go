package cron

import (
	"github.com/robfig/cron/v3"
	"github.com/uussoop/idp-go/pkg/cron/jobs"
)

var CronJob *cron.Cron

func init() {
	CronJob = cron.New(cron.WithSeconds())

	initJobs()
}

func initJobs() {
	if CronJob != nil {
		// every 7 days get models 0 0 0 ? * 7/7 *
		CronJob.AddFunc("0 0 0 28 */2 *", jobs.RefreshKeys)

	}
}

func Start() {
	if CronJob != nil {

		CronJob.Start()
	}
}
