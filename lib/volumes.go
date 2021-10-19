package lib

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

func (b *Beluga) GetAllVolumes() []*types.Volume {
	volumes, err := b.BClient.VolumeList(b.BContext, filters.NewArgs())
	if err != nil {
		panic("Error getting Docker images information")
	}

	return volumes.Volumes
}
