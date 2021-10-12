package lib

import (
	"context"

	"github.com/docker/docker/client"
)

type Beluga struct {
	BClient  client.Client
	BContext context.Context
}

func NewBeluga() *Beluga {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic("Error creating Docker client")
	}
	bel := &Beluga{
		BClient:  *cli,
		BContext: context.Background(),
	}
	return bel
}
