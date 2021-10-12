package lib

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
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

func (b *Beluga) GetAllImages() []types.ImageSummary {
	opts := types.ImageListOptions{
		All:     true,
		Filters: filters.NewArgs(),
	}
	imgs, err := b.BClient.ImageList(b.BContext, opts)
	if err != nil {
		panic("Error getting Docker images information")
	}

	return imgs
}

func (b *Beluga) GetAllContainers() []types.Container {
	opts := types.ContainerListOptions{
		Quiet:   false,
		Size:    true,
		All:     true,
		Latest:  true,
		Since:   "",
		Before:  "",
		Limit:   10,
		Filters: filters.NewArgs(),
	}
	cons, err := b.BClient.ContainerList(b.BContext, opts)
	if err != nil {
		panic("Error getting Docker containers information")
	}

	return cons
}
