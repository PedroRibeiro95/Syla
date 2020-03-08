package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PedroRibeiro95/syla/internal/handler"
	"github.com/PedroRibeiro95/syla/pkg/provider/spotify"
)

func main() {
	spotifyProvider := spotify.New("clientid", "secretkey", "http://localhost:8080/auth")
	spotifyHandler := handler.New(spotifyProvider)

	spotifyAuthHandler := handler.SpotifyAuthHandler{}

	fmt.Println(spotifyProvider.URL)

	go func() {
		// Registers Spotify Authenticator handler
		http.Handle("/auth", &spotifyAuthHandler)

		// Register API handlers
		http.HandleFunc("/api/spotify/falbums", spotifyHandler.GetFavoriteAlbumsAPI())
		http.HandleFunc("/api/spotify/fartists", spotifyHandler.GetFavoriteArtistsAPI())

		// Listens indefinetly
		fmt.Println("Listening on 8080...")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Waits for the callback...
	for spotifyAuthHandler.Request == nil {
	}

	spotifyProvider.InstantiateClient(spotifyAuthHandler.Request)
}
