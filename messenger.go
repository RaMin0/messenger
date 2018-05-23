package messenger

import (
	"net/url"

	"github.com/ramin0/request"
)

const (
	graphURL = "https://graph.facebook.com/v2.12"
)

// Messenger type
type Messenger struct {
	httpClient  *request.Client
	accessToken string
	verifyToken string
	callbacks   callbacks
}

// Options type
type Options struct {
	AccessToken string
	VerifyToken string
}

// New func
func New(o Options) *Messenger {
	httpClient := request.NewClient()

	httpClient.BaseURL = graphURL
	httpClient.DefaultParams.Add("access_token", o.AccessToken)
	httpClient.DefaultHeaders.Add("Content-Type", "application/json")

	return &Messenger{
		httpClient:  httpClient,
		accessToken: o.AccessToken,
		verifyToken: o.VerifyToken,
	}
}

// GetProfile func
func (ms *Messenger) GetProfile(userID string) (*Profile, error) {
	profileJSON, err := ms.httpClient.Get(userID, url.Values{"fields": []string{
		"first_name,last_name,profile_pic,locale,timezone,gender"}}, nil)
	if err != nil {
		return nil, err
	}
	return profileJSON.SafeDecode(&Profile{}).(*Profile), nil
}
