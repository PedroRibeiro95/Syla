package spotify

import (
	"time"

	"github.com/PedroRibeiro95/syla"
	"github.com/PedroRibeiro95/syla/pkg/provider"
)

// AlbumInformation ...
type AlbumInformation struct {
	Name        string
	Artist      string
	ReleaseDate time.Time
	Length      int
}

// ArtistInformation ...
type ArtistInformation struct {
	Name   string
	Albums []AlbumInformation
}

// MarshalToJSON ...
func (inf AlbumInformation) MarshalToJSON() ([]byte, error) {
	return provider.MarshalToJSON(inf)
}

// MarshalToJSON ...
func (inf ArtistInformation) MarshalToJSON() ([]byte, error) {
	return provider.MarshalToJSON(inf)
}

// Provider ...
type Provider struct {
	ClientID    string
	SecretKey   string
	RedirectURL string
}

// Type checking this provider
var _ syla.Provider = &Provider{}

// New ...
func New(ClientID, SecretKey, RedirectURL string) *Provider {
	return &Provider{
		ClientID:    ClientID,
		SecretKey:   SecretKey,
		RedirectURL: RedirectURL,
	}
}

// GetFavoriteAlbums ...
func (p *Provider) GetFavoriteAlbums() (syla.AlbumInformation, error) {
	return AlbumInformation{
		Name:        "pom pom",
		Artist:      "Ariel Pink",
		ReleaseDate: time.Now(),
		Length:      72,
	}, nil
}

// GetFavoriteArtists ...
func (p *Provider) GetFavoriteArtists() (syla.ArtistInformation, error) {
	var albumsList []AlbumInformation

	albumsList = append(albumsList, AlbumInformation{
		Name:        "pom pom",
		Artist:      "Ariel Pink",
		ReleaseDate: time.Now(),
		Length:      72,
	})

	albumsList = append(albumsList, AlbumInformation{
		Name:        "Dedicated to Bobby Jameson",
		Artist:      "Ariel Pink",
		ReleaseDate: time.Now(),
		Length:      66,
	})

	return ArtistInformation{
		Name:   "Ariel Pink",
		Albums: albumsList,
	}, nil
}
