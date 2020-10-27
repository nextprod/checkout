package main

import (
	"context"
	"log"

	"github.com/nextprod/sdk-go/runtime"
)

// Parameters represents request parameters.
type Parameters struct {
	SSHKey     string `json:"ssh-key"`
	Repository string `json:"repository"`
	Ref        string `json:"ref"`
	Depth      int    `json:"depth"`
	Submodules bool   `json:"submodules"`
}

// Event represents run event.
type Event struct {
	Parameters Parameters `json:"parameters"`
}

func run(ctx context.Context, event Event) (interface{}, error) {
	log.Printf("%v", event)
	return nil, nil
}

func main() { runtime.Start(run) }
