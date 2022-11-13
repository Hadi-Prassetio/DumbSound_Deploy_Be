package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	AuthRoutes(r)
	UserRoutes(r)
	ArtistRoutes(r)
	SongRoutes(r)
	TransactionRoutes(r)
}
