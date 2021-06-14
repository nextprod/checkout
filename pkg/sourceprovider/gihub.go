package sourceprovider

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
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
func (p *gitProvider) Download(ctx context.Context, privateKey []byte, url, branch, commit, path string) error {
	var signer ssh.Signer
	if privateKey != nil {
		key, _ := pem.Decode(privateKey)
		var pkey interface{}
		if key == nil {
			return fmt.Errorf("ssh: invalid key")
		}
		switch key.Type {
		case "PRIVATE KEY":
			rsa, err := x509.ParsePKCS8PrivateKey(key.Bytes)
			if err != nil {
				return err
			}
			pkey = rsa
		default:
			return fmt.Errorf("ssh: unsupported key type %q", key.Type)
		}
		var err error
		signer, err = ssh.NewSignerFromKey(pkey)
		if err != nil {
			return err
		}
	}
	options := &git.CloneOptions{
		URL:               url,
		ReferenceName:     plumbing.NewBranchReferenceName(branch),
		SingleBranch:      true,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress:          os.Stdout,
	}
	if signer != nil {
		options.Auth = &gitssh.PublicKeys{
			User:                  "git",
			Signer:                signer,
			HostKeyCallbackHelper: defaultHostKeyCallbackHelper,
		}
	}
	r, err := git.PlainCloneContext(ctx, path, false, options)
	if err != nil {
		return err
	}
	tree, err := r.Worktree()
	if err != nil {
		return err
	}
	checkoutOptions := &git.CheckoutOptions{
		Hash: plumbing.NewHash(commit),
	}
	if err := tree.Checkout(checkoutOptions); err != nil {
		return err
	}
	return nil
}
