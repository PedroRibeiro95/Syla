package spotify

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PedroRibeiro95/syla"
	"github.com/PedroRibeiro95/syla/pkg/provider"
	"github.com/zmb3/spotify"
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
	ClientID      string
	SecretKey     string
	RedirectURL   string
	Authenticator spotify.Authenticator
	URL           string
	Client        spotify.Client
}

// Type checking this provider
var _ syla.Provider = &Provider{}

// New ...
func New(clientid, secretkey, redirecturl string) *Provider {
	auth := spotify.NewAuthenticator(redirecturl, spotify.ScopeUserTopRead)
	auth.SetAuthInfo(clientid, secretkey)

	url := auth.AuthURL("test")

	return &Provider{
		ClientID:      clientid,
		SecretKey:     secretkey,
		RedirectURL:   redirecturl,
		Authenticator: auth,
		URL:           url,
	}
}

// InstantiateClient ...
func (p *Provider) InstantiateClient(r *http.Request) {
	// use the same state string here that you used to generate the URL
	token, err := p.Authenticator.Token("test", r)
	if err != nil {
		fmt.Println("Error!")
	}
	// create a client using the specified token
	p.Client = p.Authenticator.NewClient(token)

	// the client can now be used to make authenticated requests
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
