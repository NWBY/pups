package lib

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

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
