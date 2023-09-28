package datev_api_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"golang.org/x/oauth2"
)

func TestAccessTokenRevocation(t *testing.T) {
	req := client.NewTokenRevocationRequest()
	token, err := client.Oauth().TokenSource(context.Background(), &oauth2.Token{RefreshToken: os.Getenv("REFRESH_TOKEN")}).Token()
	if err != nil {
		t.Fatal(err)
	}
	req.FormParams().Token = token.AccessToken
	req.FormParams().TokenTypeHint = "access_token"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

func TestRefreshTokenRevocation(t *testing.T) {
	req := client.NewTokenRevocationRequest()
	token, err := client.Oauth().TokenSource(context.Background(), &oauth2.Token{RefreshToken: os.Getenv("REFRESH_TOKEN")}).Token()
	if err != nil {
		t.Fatal(err)
	}
	req.FormParams().Token = token.RefreshToken
	// req.FormParams().Token = os.Getenv("REFRESH_TOKEN")
	req.FormParams().TokenTypeHint = "refresh_token"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
