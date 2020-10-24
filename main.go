package main

import (
	"context"
	"log"

	"github.com/nextprod/sdk-go/runtime"
)

func run(ctx context.Context, payload []byte) (interface{}, error) {
	log.Printf("Event: %s", string(payload))
	return nil, nil
}

func main() { runtime.Start(run) }
