package syla

// AlbumInformation ...
type AlbumInformation interface {
	MarshalToJSON() ([]byte, error)
}

// ArtistInformation ...
type ArtistInformation interface {
	MarshalToJSON() ([]byte, error)
}

// Provider ...
type Provider interface {
	GetFavoriteAlbums() (AlbumInformation, error)
	GetFavoriteArtists() (ArtistInformation, error)
}
