package syla

// FavoriteAlbumsInformation ...
type FavoriteAlbumsInformation interface {
	MarshalToJSON() string
}

// FavoriteArtistsInformation ...
type FavoriteArtistsInformation interface {
	MarshalToJSON() string
}

// Provider ...
type Provider interface {
	GetFavoriteAlbums() (FavoriteAlbumsInformation, error)
	GetFavoriteArtists() (FavoriteArtistsInformation, error)
}
