package routes

import (
	"dumbsound/handlers"
	"dumbsound/pkg/middleware"
	"dumbsound/pkg/mysql"
	"dumbsound/repositories"

	"github.com/gorilla/mux"
)

func ArtistRoutes(r *mux.Router) {
	artistRepository := repositories.RepositoryArtist(mysql.DB)
	h := handlers.HandlerArtist(artistRepository)

	r.HandleFunc("/artists", middleware.Auth(h.FindArtists)).Methods("GET")
	r.HandleFunc("/artist/{id}", middleware.Auth(h.GetArtist)).Methods("GET")
	r.HandleFunc("/artist", middleware.Auth(h.CreateArtist)).Methods("POST")
}
