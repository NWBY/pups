package lib

import (
	"github.com/containers/podman/v3/pkg/bindings/images"
	"github.com/containers/podman/v3/pkg/domain/entities"
)

func (p *Pups) GetAllImages() []*entities.ImageSummary {
	images, _ := images.List(p.PConnection, &images.ListOptions{})

	return images
}
