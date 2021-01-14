package verifier

import "context"

type Verifier interface {
	Verify(ctx context.Context, token string) (userID string, err error)
}
