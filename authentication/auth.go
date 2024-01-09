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

	code := queryParts["code"][0]

	tok, err := conf.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		log.Fatal(err)
	}
	authToken = tok.AccessToken
	_, err = io.WriteString(w, "Login successful! Please close this window and return back to CLI")
	if err != nil {
		log.Fatal(err)
	}
}

func serve(done chan int) {
	log.Fatal(http.ListenAndServe(":9998", nil))
	if done != nil {
		done <- 0
	}
}

func VerifyLogin(clientID string, clientSecret string) (string, error) {
	authToken = ""
	ctx = context.Background()
	conf = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"openid"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://sso.stage.redhat.com/auth/realms/redhat-external/protocol/openid-connect/auth",
			TokenURL: "https://sso.stage.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token",
		},
		RedirectURL: "http://127.0.0.1:9998/oauth/callback",
	}
	verifier = oauth2.GenerateVerifier()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	sslcli := &http.Client{Transport: tr}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, sslcli)
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))

	http.HandleFunc("/oauth/callback", callbackHandler)
	done := make(chan int)
	twoMinTimer := time.Now().Local().Add(time.Minute * 2)
	go serve(done)
	time.Sleep(2 * time.Second)
	err := open.Run(url)
	if err != nil {
		return authToken, err
	}
	time.Sleep(1 * time.Second)
	for {
		if authToken != "" {
			return authToken, nil
		}
		if time.Now().After(twoMinTimer) {
			return authToken, fmt.Errorf("Time expired")
		}
	}
}
