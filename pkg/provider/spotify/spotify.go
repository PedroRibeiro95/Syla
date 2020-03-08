package spotify

import (
	"fmt"
	"net/http"

	"github.com/PedroRibeiro95/syla"
	"github.com/PedroRibeiro95/syla/pkg/provider"
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
	auth := spotify.NewAuthenticator(redirecturl, spotify.ScopeUserTopRead, spotify.ScopeUserFollowRead, spotify.ScopeUserLibraryRead)
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
// TODO verify that the client has been instantiated
// TODO understand pagination
func (p *Provider) GetFavoriteAlbums() ([]syla.AlbumInformation, error) {

	var favoriteAlbumsResponse []syla.AlbumInformation

	albumsList, err := p.Client.CurrentUsersAlbums()
	if err != nil {
		fmt.Println("Oh my god! Error!")
		fmt.Println(err)
		return []syla.AlbumInformation{}, nil
	}

	for albumsList.Next != "" {
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

	for followedArtists.Next != "" {
		for _, artist := range followedArtists.Artists {
			artistInformation := ArtistInformation{
				Name:           artist.Name,
				Popularity:     artist.Popularity,
				Genres:         artist.Genres,
				FollowersCount: artist.Followers.Count,
			}

			favoriteArtistsResponse = append(favoriteArtistsResponse, artistInformation)
		}

	}

	return favoriteArtistsResponse, nil
}
