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

// GenericProviderHandler ...
type GenericProviderHandler struct {
	Provider *syla.Provider
}

// New ...
func New(p syla.Provider) *GenericProviderHandler {
	return &GenericProviderHandler{
		Provider: &p,
	}
}

// GetFavoriteAlbums ...
func (ph *GenericProviderHandler) GetFavoriteAlbums() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		//No checks at this point
		resp, err := (*ph.Provider).GetFavoriteAlbums()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		jsonResp, err := resp.MarshalToJSON()
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	}
}
