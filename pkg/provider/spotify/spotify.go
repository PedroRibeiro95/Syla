package spotify

import (
	"fmt"
	"net/http"

	"github.com/PedroRibeiro95/syla"
	"github.com/PedroRibeiro95/syla/pkg/provider"
	log "github.com/sirupsen/logrus"
	"github.com/zmb3/spotify"
)

// FavoriteArtistsResponse ...
type FavoriteArtistsResponse struct {
	Next            string
	FavoriteArtists []ArtistInformation
}

// FavoriteAlbumsResponse ...
type FavoriteAlbumsResponse struct {
	Limit          int
	Offset         int
	FavoriteAlbums []AlbumInformation
}

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
func (inf FavoriteAlbumsResponse) MarshalToJSON() ([]byte, error) {
	return provider.MarshalToJSON(inf)
}

// MarshalToJSON ...
func (inf FavoriteArtistsResponse) MarshalToJSON() ([]byte, error) {
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
func (p *Provider) GetFavoriteAlbums(limit, offset int) (syla.FavoriteAlbumsResponse, error) {

	var favoriteAlbums []AlbumInformation

	albumsList, err := p.Client.CurrentUsersAlbumsOpt(&spotify.Options{
		Limit:  &limit,
		Offset: &offset,
	})
	if err != nil {
		log.Error("Oh my god! Error!")
		log.Error(err)
		return FavoriteAlbumsResponse{}, nil
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

		favoriteAlbums = append(favoriteAlbums, albumInformation)
	}

	return FavoriteAlbumsResponse{
		Limit:          albumsList.Limit,
		Offset:         albumsList.Offset,
		FavoriteAlbums: favoriteAlbums,
	}, nil
}

// GetFavoriteArtists ...
func (p *Provider) GetFavoriteArtists(limit int, next string) (syla.FavoriteArtistsResponse, error) {
	var favoriteArtists []ArtistInformation

	followedArtists, err := p.Client.CurrentUsersFollowedArtistsOpt(limit, next)
	if err != nil {
		log.Error("Oh my god! Error!")
		log.Error(err)
		return FavoriteArtistsResponse{}, nil
	}

	for _, artist := range followedArtists.Artists {
		artistInformation := ArtistInformation{
			Name:           artist.Name,
			Popularity:     artist.Popularity,
			Genres:         artist.Genres,
			FollowersCount: artist.Followers.Count,
		}

		favoriteArtists = append(favoriteArtists, artistInformation)
	}

	return FavoriteArtistsResponse{
		Next:            followedArtists.Next,
		FavoriteArtists: favoriteArtists,
	}, nil
}
