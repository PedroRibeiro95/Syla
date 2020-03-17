package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PedroRibeiro95/syla"
	log "github.com/sirupsen/logrus"
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

// SpotifyAuthHandler ...
type SpotifyAuthHandler struct {
	Request *http.Request
}

var _ http.Handler = &SpotifyAuthHandler{}

// Verifies that GenericProviderHandler implements the ProviderHandler interface
var _ ProviderHandler = &GenericProviderHandler{}

// New ...
func New(p syla.Provider) *GenericProviderHandler {
	return &GenericProviderHandler{
		Provider: &p,
	}
}

// MarshalToJSON ...
func MarshalToJSON(i interface{}) ([]byte, error) {
	jsonMarshalled, err := json.Marshal(i)
	if err != nil {
		return []byte{}, err
	}
	return jsonMarshalled, nil
}

// GetFavoriteAlbumsAPI ...
func (ph *GenericProviderHandler) GetFavoriteAlbumsAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		log.Info("Getting request for Favorite Albums")
		//No checks at this point
		response, err := (*ph.Provider).GetFavoriteAlbums()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		jsonResponse, err := MarshalToJSON(response)
		if err != nil {
			fmt.Println("Hory shet, an error!")
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}

// GetFavoriteArtistsAPI ...
func (ph *GenericProviderHandler) GetFavoriteArtistsAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		log.Info("Getting request for Favorite Artists")
		//No checks at this point
		response, err := (*ph.Provider).GetFavoriteArtists()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		jsonResponse, err := MarshalToJSON(response)
		if err != nil {
			fmt.Println("Hory shet, an error!")
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}

// ServeHTTP ...
func (sh *SpotifyAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sh.Request = r
}
