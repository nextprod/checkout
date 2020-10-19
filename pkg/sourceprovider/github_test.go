package sourceprovider

import (
	"context"
	"testing"
)

func TestGithubCheckout(t *testing.T) {
	p := NewGitProvider()
	if err := p.Download(context.Background(), "https://github.com/nextprod/checkout", "/home/justin/go/src/github.com/nextprod/repos/checkout"); err != nil {
		t.Fatal(err)
	}
}
