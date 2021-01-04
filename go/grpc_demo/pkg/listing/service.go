package listing

import "errors"

// ErrNotFound is used when a movie could not be found for a specified
var ErrNotFound = errors.New("movie not found")

// Repository provides access to the movie and castmembers storage
type Repository interface {
	// GetMovie returns the movie with specified ID
	GetMovie(ID int) (Movie, error)
	// GetAllMovies returns all movies saved in storage
	GetAllMovies() ([]Movie, error)
}

// Service provides beer and review listing operations
type Service interface {
	GetMovie(id int) (Movie, error)
	GetMovies() ([]Movie, error)
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetMovies returns all movies
func (svc *service) GetMovies() ([]Movie, error) {
	return svc.r.GetAllMovies()
}

// GetMovie returns the movie with the specified ID
func (svc *service) GetMovie(ID int) (Movie, error) {
	return svc.r.GetMovie(ID)
}
