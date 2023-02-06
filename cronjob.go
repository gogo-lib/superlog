package superlog

import (
	"github.com/robfig/cron/v3"
)

var _cronjob *cron.Cron

func init() {
	_cronjob = cron.New()
	_cronjob.Start()
}
