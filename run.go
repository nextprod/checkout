package main

import (
	"context"
	"encoding/base64"
	"os"

	"github.com/nextprod/checkout/pkg/sourceprovider"
)

// Parameters represents request parameters.
type Parameters struct {
	SSHKey     *string `json:"ssh-key"`
	Repository string  `json:"repository"`
	Branch     string  `json:"branch"`
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
	if err := provider.Download(ctx, keyb, params.Repository, params.Branch, params.Ref, os.Getenv("NEX_WORKSPACE")); err != nil {
		return nil, err
	}
	return nil, nil
}
