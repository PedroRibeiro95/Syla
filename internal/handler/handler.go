package handler

import (
	"net/http"

	"github.com/PedroRibeiro95/syla"
)

// ProviderHandler ...
type ProviderHandler interface {
	GetFavoriteAlbumsAPI() http.HandlerFunc
	GetFavoriteArtistsAPI() http.HandlerFunc
}

// GenericProviderHandler ...
type GenericProviderHandler struct {
	Provider *syla.Provider
}

// Verifies that GenericProviderHandler implements the ProviderHandler interface
var _ ProviderHandler = &GenericProviderHandler{}

// New ...
func New(p syla.Provider) *GenericProviderHandler {
	return &GenericProviderHandler{
		Provider: &p,
	}
}

// GetFavoriteAlbumsAPI ...
func (ph *GenericProviderHandler) GetFavoriteAlbumsAPI() http.HandlerFunc {
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

// GetFavoriteArtistsAPI ...
func (ph *GenericProviderHandler) GetFavoriteArtistsAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		//No checks at this point
		resp, err := (*ph.Provider).GetFavoriteArtists()
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
