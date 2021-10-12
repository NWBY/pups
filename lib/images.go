package lib

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

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
