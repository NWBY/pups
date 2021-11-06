package lib

import (
	"github.com/containers/podman/v3/pkg/bindings/containers"
	"github.com/containers/podman/v3/pkg/domain/entities"
)

func (p *Pups) GetAllContainers() []entities.ListContainer {
	containers, _ := containers.List(p.PConnection, &containers.ListOptions{})

	return containers
}
