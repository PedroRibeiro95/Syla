package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PedroRibeiro95/syla/internal/handler"
	"github.com/PedroRibeiro95/syla/pkg/provider/spotify"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a test of HTTP Handlers: %s", r.URL.Path)
}

func main() {
	spotifyProvider := spotify.New("clientid", "secretkey", "redirecturl")
	spotifyHandler := handler.New(spotifyProvider)

	http.HandleFunc("/spotify/favalbums", spotifyHandler.GetFavoriteAlbums())

	// Listens indefinetly
	log.Fatal(http.ListenAndServe(":8080", nil))
}
