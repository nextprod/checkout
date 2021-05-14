package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
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
	os.Stdout.WriteString("Reading input...\n")
	reader := bufio.NewReader(os.Stdin)
	in, err := reader.ReadString('\n')
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}
	var event Event
	if err := json.Unmarshal([]byte(in), &event); err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}
	if _, err := run(context.Background(), event); err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}
}
