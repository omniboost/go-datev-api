package datev_api_test

import (
	"context"
	"os"
	"testing"
	"time"

	datev "github.com/omniboost/go-datev-api"
	"golang.org/x/oauth2"
)

var (
	client     *datev.Client
	businessID int
)

func TestMain(m *testing.M) {
	baseURL := os.Getenv("BASE_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	refreshToken := os.Getenv("REFRESH_TOKEN")
	accessToken := os.Getenv("ACCESS_TOKEN")
	datevClientID := os.Getenv("DATEV_CLIENT_ID")
	tokenURL := os.Getenv("TOKEN_URL")
	revokeURL := os.Getenv("REVOKE_URL")
	debug := os.Getenv("DEBUG")

	oauthConfig := datev.NewOauth2Config()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.Endpoint.TokenURL = tokenURL
	}

	token := &oauth2.Token{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
		Expiry:       time.Now().AddDate(0, 0, 1),
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(context.Background(), token)

	client = datev.NewClient(httpClient)
	client.SetClientID(clientID)
	client.SetClientSecret(clientSecret)
	client.SetDatevClientID(datevClientID)
	client.SetOauth(oauthConfig)

	if debug != "" {
		client.SetDebug(true)
	}

	if revokeURL != "" {
		client.SetRevokeURL(revokeURL)
	}

	if baseURL != "" {
		client.SetBaseURL(baseURL)
	}

	client.SetDisallowUnknownFields(true)
	m.Run()
}
