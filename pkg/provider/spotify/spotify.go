package spotify

import (
	"fmt"
	"net/http"

	"github.com/PedroRibeiro95/syla"
	"github.com/PedroRibeiro95/syla/pkg/provider"
	log "github.com/sirupsen/logrus"
	"github.com/zmb3/spotify"
)

// AlbumInformation ...
type AlbumInformation struct {
	Name        string
	Artists     []string
	ReleaseDate string
	URLs        map[string]string
	Genres      []string
}

// ArtistInformation ...
type ArtistInformation struct {
	Name           string
	Popularity     int
	Genres         []string
	FollowersCount uint
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
	Authenticator spotify.Authenticator
	URL           string
	Client        spotify.Client
	Read          bool
}

// Type checking this provider
var _ syla.Provider = &Provider{}

// New ...
func New(clientid, secretkey, redirecturl string) *Provider {
	p := Provider{}

	p.Authenticator = spotify.NewAuthenticator(redirecturl, spotify.ScopeUserTopRead, spotify.ScopeUserFollowRead, spotify.ScopeUserLibraryRead)
	p.Authenticator.SetAuthInfo(clientid, secretkey)
	p.URL = p.Authenticator.AuthURL("syla")

	return &p
}

// InstantiateClient ...
func (p *Provider) InstantiateClient(r *http.Request) {
	log.Info("Got callback!")
	// use the same state string here that you used to generate the URL
	token, err := p.Authenticator.Token("syla", r)
	if err != nil {
		fmt.Println("Error!")
	}
	log.Debug("Got token from request")
	// create a client using the specified token
	p.Client = p.Authenticator.NewClient(token)
	// the client can now be used to make authenticated requests
}

// GetFavoriteAlbums ...
// TODO verify that the client has been instantiated
// TODO understand pagination
func (p *Provider) GetFavoriteAlbums() ([]syla.AlbumInformation, error) {

	var favoriteAlbumsResponse []syla.AlbumInformation

	albumsList, err := p.Client.CurrentUsersAlbums()
	if err != nil {
		log.Error("Oh my god! Error!")
		fmt.Println(err)
		return []syla.AlbumInformation{}, nil
	}

	for _, album := range albumsList.Albums {
		albumInformation := AlbumInformation{
			Name:        album.Name,
			ReleaseDate: album.ReleaseDate,
			URLs:        album.ExternalURLs,
			Genres:      album.Genres,
		}

		var artists []string
		for _, artist := range album.Artists {
			artists = append(artists, artist.Name)
		}
		albumInformation.Artists = artists

		favoriteAlbumsResponse = append(favoriteAlbumsResponse, albumInformation)
	}

	return favoriteAlbumsResponse, nil
}

// GetFavoriteArtists ...
func (p *Provider) GetFavoriteArtists() ([]syla.ArtistInformation, error) {
	var favoriteArtistsResponse []syla.ArtistInformation

	followedArtists, err := p.Client.CurrentUsersFollowedArtists()
	if err != nil {
		fmt.Println("Oh my god! Error!")
		return []syla.ArtistInformation{}, nil
	}

	for _, artist := range followedArtists.Artists {
		artistInformation := ArtistInformation{
			Name:           artist.Name,
			Popularity:     artist.Popularity,
			Genres:         artist.Genres,
			FollowersCount: artist.Followers.Count,
		}

		favoriteArtistsResponse = append(favoriteArtistsResponse, artistInformation)
	}

	return favoriteArtistsResponse, nil
}
