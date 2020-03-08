package syla

// FavoriteAlbumsInformation ...
type FavoriteAlbumsInformation interface {
	MarshalToJSON() ([]byte, error)
}

// FavoriteArtistsInformation ...
type FavoriteArtistsInformation interface {
	MarshalToJSON() ([]byte, error)
}

// Provider ...
type Provider interface {
	GetFavoriteAlbums() (FavoriteAlbumsInformation, error)
	GetFavoriteArtists() (FavoriteArtistsInformation, error)
}
