package common

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

// Starts the cron job.
func CronStart(intervals [] uint64, jobs ...func()) {
	fmt.Println("Creating cron jobs...")

	for index, job := range jobs {
		gocron.Every(intervals[index]).Hours().Do(job)
	}

	fmt.Println("Cron jobs created...")
	<-gocron.Start()
}
