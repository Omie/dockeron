package dockeron

import (
	"time"
)

func worker(job *Job) {
	containerId, err := dk.MakeContainer(job)
	if err != nil {
		return
	}

	for {
		err = dk.StartContainer(containerId, job)
		if err != nil {
			return
		}
		time.Sleep(time.Duration(job.Interval) * time.Minute)
	}

}
