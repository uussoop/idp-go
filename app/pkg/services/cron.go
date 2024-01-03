package services

import "github.com/uussoop/idp-go/pkg/cron"

func InitCron() {
	cron.Start()
}
