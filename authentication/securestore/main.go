package securestore

const (
	LabelKey             = "Red Hat SSO"
	SecureStoreConfigKey = "securestore"
)

// type Credentials struct {
// 	AccessToken  string   `json:"access_token,omitempty" doc:"Bearer access token."`
// 	ClientID     string   `json:"client_id,omitempty" doc:"OpenID client identifier."`
// 	ClientSecret string   `json:"client_secret,omitempty" doc:"OpenID client secret."`
// 	Insecure     bool     `json:"insecure,omitempty" doc:"Enables insecure communication with the server. This disables verification of TLS certificates and host names."`
// 	Password     string   `json:"password,omitempty" doc:"User password."`
// 	RefreshToken string   `json:"refresh_token,omitempty" doc:"Offline or refresh token."`
// 	Scopes       []string `json:"scopes,omitempty" doc:"OpenID scope. If this option is used it will replace completely the default scopes. Can be repeated multiple times to specify multiple scopes."`
// 	TokenURL     string   `json:"token_url,omitempty" doc:"OpenID token URL."`
// 	URL          string   `json:"url,omitempty" doc:"URL of the API gateway. The value can be the complete URL or an alias. The valid aliases are 'production', 'staging' and 'integration'."`
// 	User         string   `json:"user,omitempty" doc:"User name."`
// 	Pager        string   `json:"pager,omitempty" doc:"Pager command, for example 'less'. If empty no pager will be used."`
// }

func UpsertConfigToKeyring(credentials []byte) error {
	return upsertCredentials(LabelKey, credentials)
}

func GetConfigFromKeyring() ([]byte, error) {
	return getCredentials(LabelKey)
}
