package spotify

import (
	"github.com/PedroRibeiro95/syla"
	"github.com/PedroRibeiro95/syla/pkg/provider"
)

// FavoriteAlbumsInformation ...
type FavoriteAlbumsInformation struct {
	Name string
}

// MarshalToJSON ...
func (inf FavoriteAlbumsInformation) MarshalToJSON() ([]byte, error) {
	return provider.MarshalToJSON(inf)
}

// Provider ...
type Provider struct {
	ClientID    string
	SecretKey   string
	RedirectURL string
}

// Type checking this provider
var _ syla.Provider = Provider{}

// New ...
func New(ClientID, SecretKey, RedirectURL string) *Provider {
	return &Provider{
		ClientID:    ClientID,
		SecretKey:   SecretKey,
		RedirectURL: RedirectURL,
	}
}

// GetFavoriteAlbums ...
func (p *Provider) GetFavoriteAlbums() (syla.FavoriteAlbumsInformation, error) {
	return FavoriteAlbumsInformation{
		Name: "Test!",
	}, nil
}
