package dockeron

import (
	"github.com/samalba/dockerclient"
)

var (
	dk *Dockeron
)

func init() {
	// try to connect to docker socket,
	// if it fails, makes no sense to go any further, just panic
	docker, err := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
	if err != nil {
		panic(err)
	}
	dk = &Dockeron{client: docker}
}

func StartJobs(jobs []*Job) {
	for _, j := range jobs {
		go worker(j)
	}
}
