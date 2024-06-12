package utils

import (
	"os"

	log "observa/internal/pkg/log"

	"github.com/robfig/cron/v3"
)

func ScheduleAndStart(spec string, cmd func()) {
	var job = Schedule(spec, cmd)
	log.Info("Starting the job")
	job.Start()
}
func Schedule(spec string, cmd func()) *cron.Cron {
	c := cron.New()
	_, err := c.AddFunc(spec, cmd)
	if err != nil {
		log.Error("Error scheduling job:", err)
		os.Exit(1)
	}
	return c
}
