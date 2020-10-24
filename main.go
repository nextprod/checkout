package main

import (
	"context"
	"log"

	"github.com/nextprod/sdk-go/runtime"
)

func run(ctx context.Context, payload []byte) (interface{}, error) {
	log.Println("Received request")
	return nil, nil
}

func main() {
	runtime.Start(run)
}