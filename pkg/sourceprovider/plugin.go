package sourceprovider

import (
	"context"
)

// SourceProvider represents code source provider.
type SourceProvider interface {
	// Download downloads source from source provider.
	Download(ctx context.Context, ssh []byte, url, ref, path string) error
}
