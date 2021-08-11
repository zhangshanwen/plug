package docker

import (
	"github.com/docker/docker/client"
)

var (
	Docker *client.Client
)

func InitDocker() (err error) {
	Docker, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	return err
}
