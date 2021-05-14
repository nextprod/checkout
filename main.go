package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/nextprod/checkout/pkg/sourceprovider"
)

// Parameters represents request parameters.
type Parameters struct {
	SSHKey     *string `json:"ssh-key"`
	Repository string  `json:"repository"`
	Ref        string  `json:"ref"`
	Depth      string  `json:"depth"`
	Submodules string  `json:"submodules"`
}

// Event represents run event.
type Event struct {
	Parameters Parameters `json:"parameters"`
}

func run(ctx context.Context, event Event) (interface{}, error) {
	params := event.Parameters
	var keyb []byte
	if params.SSHKey != nil {
		var err error
		// Decode base64 encoded SSH private key.
		keyb, err = base64.StdEncoding.DecodeString(*params.SSHKey)
		if err != nil {
			return nil, err
		}
	}
	provider := sourceprovider.NewGitProvider()
	if err := provider.Download(ctx, keyb, params.Repository, params.Ref, os.Getenv("NEX_WORKSPACE")); err != nil {
		return nil, err
	}
	return nil, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	in, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	var event Event
	if err := json.Unmarshal([]byte(in), &event); err != nil {
		panic(err)
	}
	if _, err := run(context.Background(), event); err != nil {
		panic(err)
	}
	fmt.Println("Parameters parsed")
}
