package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/nextprod/checkout/pkg/sourceprovider"
)

// Parameters represents request parameters.
type Parameters struct {
	SSHKey     *string `json:"ssh-key"`
	Repository string  `json:"repository"`
	Ref        string  `json:"ref"`
	Depth      int     `json:"depth"`
	Submodules bool    `json:"submodules"`
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
	event, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(event)
}
