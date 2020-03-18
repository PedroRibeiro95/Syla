package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PedroRibeiro95/syla"
	"github.com/gorilla/mux"
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

func (ph *GenericProviderHandler) getLimitAndOffset(r *http.Request) (int, int, error) {
	var (
		limit, offset int64
		err           error
	)

	vars := mux.Vars(r)

	limitString := vars["limit"]
	offsetString := vars["offset"]

	limit, err = strconv.ParseInt(limitString, 10, 0)
	if err != nil {
		log.Warn("Error converting limit to integer")
		return 0, 0, err
	}
	offset, err = strconv.ParseInt(offsetString, 10, 0)
	if err != nil {
		log.Warn("Error converting offset to integer")
		return 0, 0, err
	}

	if limit == 0 {
		return 20, int(offset), nil
	}

	return int(limit), int(offset), nil

}

// GetFavoriteAlbumsAPI ...
func (ph *GenericProviderHandler) GetFavoriteAlbumsAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			limit, offset int64
			err           error
			response      syla.FavoriteAlbumsResponse
		)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		log.Info("Getting request for Favorite Albums")

		//No checks at this point
		log.Debug("Getting 'limit' and 'next' for this request")
		vars := mux.Vars(r)

		limitString := vars["limit"]
		offsetString := vars["offset"]

		limit, err = strconv.ParseInt(limitString, 10, 64)
		if err != nil {
			log.Warn("Error converting limit to integer")
		}
		offset, err = strconv.ParseInt(offsetString, 10, 64)
		if err != nil {
			log.Warn("Error converting offset to integer")
		}

		if limit == 0 {
			response, err = (*ph.Provider).GetFavoriteAlbums(20, int(offset))
		} else {
			response, err = (*ph.Provider).GetFavoriteAlbums(int(limit), int(offset))
		}
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
		var (
			limit    int64
			next     string
			err      error
			response syla.FavoriteArtistsResponse
		)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		log.Info("Getting request for Favorite Artists")

		log.Debug("Getting 'limit' and 'next' for this request")
		vars := mux.Vars(r)

		limitString := vars["limit"]
		next = vars["next"]

		limit, err = strconv.ParseInt(limitString, 10, 0)
		if err != nil {
			log.Warn("Error converting limit to integer")
		}

		if next == "nil" {
			next = ""
		}

		//No checks at this point
		if limit == 0 {
			response, err = (*ph.Provider).GetFavoriteArtists(10, next)
		} else {
			response, err = (*ph.Provider).GetFavoriteArtists(int(limit), next)
		}
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
