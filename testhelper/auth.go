package testhelper

import (
	"context"
	"errors"
)

const ValidAuthJWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.4TFMMEj2Ejof4za_5H_CVsM2PQX3YDYIMOZ4t0LcWjA"

type AuthVerifier struct{}

func (tv *AuthVerifier) Verify(ctx context.Context, idToken string) (string, error) {
	if idToken != ValidAuthJWT {
		return "", errors.New("invalid token")
	}
	return "testUser", nil
}
