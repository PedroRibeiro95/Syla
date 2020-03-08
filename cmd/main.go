package main

import (
	"fmt"
	"log"
	"net/http"

	"syla/handler"
	"syla/spotify"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a test of HTTP Handlers: %s", r.URL.Path)
}

func main() {
	//http.HandleFunc("/", test)

	spotifyProvider := spotify.New("clientid", "secretkey", "redirecturl")
	spotifyHandler := handler.GenericProviderHandler(spotifyProvider)

	http.HandleFunc("/spotify/favalbums", spotifyHandler.GetFavoriteAlbums())

	// Listens indefinetly
	log.Fatal(http.ListenAndServe(":8080", nil))
}
