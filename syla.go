package syla

// FavoriteAlbumsResponse ...
type FavoriteAlbumsResponse interface {
	MarshalToJSON() ([]byte, error)
}

// FavoriteArtistsResponse ...
type FavoriteArtistsResponse interface {
	MarshalToJSON() ([]byte, error)
}

// Provider ...
type Provider interface {
	GetFavoriteAlbums(int, int) (FavoriteAlbumsResponse, error)
	GetFavoriteArtists(int, string) (FavoriteArtistsResponse, error)
}
