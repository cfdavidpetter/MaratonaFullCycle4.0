package main

import (
	"context"
	"log"

	oidc "github.com/coreos/go-oidc"
)

var (
	clientID     = "app"
	clientSecret = "790e27c6-457f-4388-8812-effdb5a52505"
)

func main() {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/auth/realms/demo")
	if err != nil {
		log.Fatal(err)
	}

	config := oauth2.Config{
		clientID:     clientID,
		clientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:8081/auth/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

}
