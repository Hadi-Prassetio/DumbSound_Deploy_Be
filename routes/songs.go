package routes

import (
	"dumbsound/handlers"
	"dumbsound/pkg/middleware"
	"dumbsound/pkg/mysql"
	"dumbsound/repositories"

	"github.com/gorilla/mux"
)

func SongRoutes(r *mux.Router) {
	songRepository := repositories.RepositorySong(mysql.DB)
	h := handlers.HandlerSong(songRepository)

	r.HandleFunc("/songs", h.FindSong).Methods("GET")
	r.HandleFunc("/song/{id}", middleware.Auth(h.GetSong)).Methods("GET")
	r.HandleFunc("/song", middleware.Auth(middleware.UploadFile(middleware.UploadSong(h.CreateSong)))).Methods("POST")
}
