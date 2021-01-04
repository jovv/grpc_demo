package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jovv/grpc_demo/go/grpc_demo/pkg/listing"
	"github.com/julienschmidt/httprouter"
)

// Handler defines routes of the REST API
func Handler(l listing.Service) http.Handler {
	router := httprouter.New()

	router.GET("/movies", getMovies(l))
	router.GET("/movie/:id", getMovie(l))

	return router
}

// getMovies returns a handler for GET /movies requests
func getMovies(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list, err := s.GetMovies()
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		json.NewEncoder(w).Encode(list)
	}

}

// getMovie returns a handler for GET /movie/:id requests
func getMovie(s listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		id, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(400), 400)
			return
		}
		movie, err := s.GetMovie(id)
		if err == listing.ErrNotFound {
			http.Error(w, "The movie you requested does not exist.", http.StatusNotFound)
			return
		}
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		json.NewEncoder(w).Encode(movie)
	}

}
