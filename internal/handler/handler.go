package handler

import (
	"net/http"
	"github.com/PedroRibeiro95/syla"
)

// ProviderHandler ...
type ProviderHandler interface {
	GetFavoriteAlbums() http.HandlerFunc
	GetFavoriteArtits() http.HandlerFunc
}
´
// GenericProviderHandler ...
type GenericProviderHandler struct {
	Provider syla.Provider
}

// New ...
func New(p syla.Provider) *GenericProviderHandler {
	return &GenericProviderHandler{
		Provider: p,
	}
}

// GetFavoriteAlbums ...
func (GenericProviderHandler ph) GetFavoriteAlbums() http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Set("Content-Type", "application/json; charset=UTF-8")

		//No checks at this point
		Provider.GetFavoriteAlbums()
	}
}
