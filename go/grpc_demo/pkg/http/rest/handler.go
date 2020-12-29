package rest

import (
	"encoding/json"
	"net/http"

	"github.com/jovv/grpc_demo/go/grpc_demo/pkg/listing"
	"github.com/julienschmidt/httprouter"
)

// Handler defines routes of the REST API
func Handler(l listing.Service) http.Handler {
	router := httprouter.New()

	router.GET("/movies", getMovies(l))
	// router.GET("/movie/:id", getMovie(l))
	// router.GET("/movie/:id/castmembers", getMovieCastMembers(l))

	return router
}

// getMovies returns a handler for GET /movies requests
func getMovies(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list := s.GetMovies()
		json.NewEncoder(w).Encode(list)
	}

}
