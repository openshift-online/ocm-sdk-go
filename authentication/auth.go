package authentication

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

var (
	conf      *oauth2.Config
	ctx       context.Context
	verifier  string
	authToken string
)

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	queryParts, _ := url.ParseQuery(r.URL.RawQuery)

	// Use the authorization code that is pushed to the redirect URL
	code := queryParts["code"][0]

	// Exchange will do the handshake to retrieve the initial access token.
	tok, err := conf.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		log.Fatal(err)
	}

	// Get the access token and ask user to go back to CLI
	authToken = tok.AccessToken
	_, err = io.WriteString(w, "Login successful! Please close this window and return back to CLI")
	if err != nil {
		log.Fatal(err)
	}
}

func serve(done chan int) {
	http.ListenAndServe(":9998", nil)
	if done != nil {
		done <- 0
	}
}

func VerifyLogin(clientID string) (string, error) {
	authToken = ""
	ctx = context.Background()
	// Create config for OAuth2, redirect to localhost for callback verification and retrieving tokens
	conf = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: "",
		Scopes:       []string{"openid"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://sso.stage.redhat.com/auth/realms/redhat-external/protocol/openid-connect/auth",
			TokenURL: "https://sso.stage.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token",
		},
		RedirectURL: "http://127.0.0.1:9998/oauth/callback",
	}
	verifier = oauth2.GenerateVerifier()

	// add transport for self-signed certificate to context
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	sslcli := &http.Client{Transport: tr}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, sslcli)

	// Create URL with PKCE
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))

	http.HandleFunc("/oauth/callback", callbackHandler)
	done := make(chan int)

	// Add a 5 min timer until the user finishes the auth process
	fiveMinTimer := time.Now().Local().Add(time.Minute * 5)
	go serve(done)
	time.Sleep(2 * time.Second)

	// Redirect user to Red Hat SSO auth page
	err := open.Run(url)
	if err != nil {
		return authToken, err
	}
	time.Sleep(1 * time.Second)
	// Wait for the user to finish auth process, and return back with authToken. Otherwise, return an error after 5 mins
	for {
		if authToken != "" {
			return authToken, nil
		}
		if time.Now().After(fiveMinTimer) {
			return authToken, fmt.Errorf("Time expired")
		}
	}
}
