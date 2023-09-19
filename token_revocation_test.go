package datev_api_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestTokenRevocation(t *testing.T) {
	req := client.NewTokenRevocationRequest()
	req.RequestBody().Token = os.Getenv("REFRESH_TOKEN")
	req.RequestBody().TokenTypeHint = "access_token"

	// req.FormParams().ClientID = os.Getenv("CLIENT_ID")
	// req.FormParams().ClientSecret = os.Getenv("CLIENT_SECRET")
	// req.FormParams().Token = os.Getenv("REFRESH_TOKEN")
	// req.FormParams().Token = "FML"
	// req.FormParams().TokenTypeHint = "access_token"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

