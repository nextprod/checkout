package sourceprovider

import (
	"context"

	"golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4"
	gitssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

// TypeGit ...
const TypeGit = "git"

// DefaultHostKeyCallbackHelper ...
var defaultHostKeyCallbackHelper = gitssh.HostKeyCallbackHelper{
	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
}

type gitProvider struct {
	//signerGetter *signerGetter
}

var _ SourceProvider = &gitProvider{}

// NewGitProvider clone github repositories into working tree.
func NewGitProvider() SourceProvider {
	p := &gitProvider{
		//signerGetter: &signerGetter{},
	}
	return p
}

// Download implements SourceProvider.
func (p *gitProvider) Download(ctx context.Context, url, path string) error {
	options := &git.CloneOptions{
		URL: url,
		Auth: &gitssh.PublicKeys{
			Signer:                nil,
			HostKeyCallbackHelper: defaultHostKeyCallbackHelper,
		},
		//Auth:              auth,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	}
	_, err := git.PlainCloneContext(ctx, path, false, options)
	return err
}
