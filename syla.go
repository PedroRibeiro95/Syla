package syla

type FavoriteAlbumsInformation interface {
	Test() string
}

type FavoriteArtistsInformation interface {
	Test() string
}

type Provider interface {
	GetFavoriteAlbums() (FavoriteAlbumsInformation, error)
	GetFavoriteArtists() (FavoriteArtistsInformation, error)
}
