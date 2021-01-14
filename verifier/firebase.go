package verifier

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
)

type firebaseVerifier struct {
	app *firebase.App
}

func NewFirebase() Verifier {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}
	return &firebaseVerifier{app}
}

func (f *firebaseVerifier) Verify(ctx context.Context, token string) (string, error) {
	client, err := f.app.Auth(ctx)
	if err != nil {
		return "", err
	}

	authToken, err := client.VerifyIDToken(ctx, token)
	if err != nil {
		return "", err
	}

	return authToken.UID, nil
}
