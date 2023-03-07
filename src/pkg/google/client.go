package google

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	utils "github.com/lokesh-go/youtube-data-golang/src/utils"
)

// Methods ...
type Methods interface {
	GetClient() (*http.Client, error)
}

type config struct {
	clientSecretPath string
	tokenPath        string
	scope            string
}

// NewClient ...
func New(clientSecretPath, tokenPath, scope string) Methods {
	// Returns
	return &config{clientSecretPath, tokenPath, scope}
}

func (c *config) GetClient() (client *http.Client, err error) {
	// Gets oauth config
	oauthConfig, err := c.getOAuthConfig()
	if err != nil {
		return nil, err
	}

	// Gets client
	client, err = c.getClient(oauthConfig)
	if err != nil {
		return nil, err
	}

	// Returns
	return client, nil
}

func (c *config) getOAuthConfig() (oauthConfig *oauth2.Config, err error) {
	// Reads credential file
	bytes, err := ioutil.ReadFile(c.clientSecretPath)
	if err != nil {
		return nil, err
	}

	// Gets google oauth config
	oauthConfig, err = google.ConfigFromJSON(bytes, c.scope)
	if err != nil {
		return nil, err
	}

	// Returns
	return oauthConfig, nil
}

func (c *config) getClient(oauthConfig *oauth2.Config) (client *http.Client, err error) {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	oauthToken, err := c.tokenFromFile()
	if err != nil {
		// Request a token from the web, then returns the retrieved token.
		oauthToken, err = getTokenFromWeb(oauthConfig)
		if err != nil {
			return nil, err
		}

		// Saves token
		err = c.saveToken(oauthToken)
		if err != nil {
			return nil, err
		}
	}

	// Refresh token if token has expired
	if oauthToken.Expiry.Before(time.Now()) {
		// Gets new token
		oauthToken, err = oauthConfig.TokenSource(context.Background(), oauthToken).Token()
		if err != nil {
			return nil, err
		}

		// Saves token
		err = c.saveToken(oauthToken)
		if err != nil {
			return nil, err
		}
	}

	// Returns
	return oauthConfig.Client(context.Background(), oauthToken), nil
}

func (c *config) tokenFromFile() (oauthToken *oauth2.Token, err error) {
	// Reads token file
	err = utils.ReadJSONFile(c.tokenPath, &oauthToken)
	if err != nil {
		return nil, err
	}

	// Returns
	return oauthToken, nil
}

func getTokenFromWeb(oauthConfig *oauth2.Config) (oauthToken *oauth2.Token, err error) {
	// Tokens from web
	authURL := oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, err
	}

	oauthToken, err = oauthConfig.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, err
	}

	// Returns
	return oauthToken, nil
}

func (c *config) saveToken(oauthToken *oauth2.Token) (err error) {
	// Saves token to the path
	file, err := os.OpenFile(c.tokenPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	json.NewEncoder(file).Encode(oauthToken)

	// Returns
	return nil
}
