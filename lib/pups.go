package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/podman/v3/pkg/bindings"
)

type Pups struct {
	PConnection context.Context
}

func NewPups() *Pups {
	conn, err := bindings.NewConnection(context.Background(), "unix://run/podman/podman.sock")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pup := &Pups{
		PConnection: conn,
	}
	return pup
}
