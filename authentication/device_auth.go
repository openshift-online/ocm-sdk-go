package authentication

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

const (
	DeviceAuthURL = "https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/auth/device"
)

// Initiates device code flow and returns access token
func InitiateDeviceAuth(clientID string) (string, error) {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: "",
		Scopes:       []string{"openid"},
		Endpoint: oauth2.Endpoint{
			DeviceAuthURL: DeviceAuthURL,
			TokenURL:      DefaultTokenURL,
		},
	}

	// Verifiers and Challenges are required for device auth
	verifier := oauth2.GenerateVerifier()
	verifierOpt := oauth2.VerifierOption(verifier)
	challenge := oauth2.S256ChallengeOption(verifier)

	// Get device code
	deviceAuthResp, err := conf.DeviceAuth(ctx, challenge, verifierOpt)
	if err != nil {
		return "", fmt.Errorf("failed to get device code: %v", err)
	}

	fmt.Printf("To continue login, navigate to %v and enter code %v\n", deviceAuthResp.VerificationURI, deviceAuthResp.UserCode)
	fmt.Printf("Checking status every %v seconds...\n", deviceAuthResp.Interval)

	// Wait for the user to enter the code, polls at interval specified in deviceAuthResp.Interval
	token, err := conf.DeviceAccessToken(ctx, deviceAuthResp, verifierOpt)
	if err != nil {
		return "", fmt.Errorf("error exchanging for access token: %v", err)
	}

	return token.RefreshToken, nil
}
