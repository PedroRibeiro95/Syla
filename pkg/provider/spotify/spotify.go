package spotify

import "syla"

type Provider struct {
	ClientID    string
	SecretKey   string
	RedirectURL string
}

// Type checking this provider
var _ syla.Provider = Provider{}

func New(ClientID, SecretKey, RedirectURL string) *Provider {
	return &Provider{
		ClientID:    ClientID,
		SecretKey:   SecretKey,
		RedirectURL: RedirectURL,
	}
}

func (p *Provider) Test() string {
	return "Successful test!"
}
