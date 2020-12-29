package listing

// Repository provides access to the movie and castmembers storage
type Repository interface {
	// GetMovie returns the movie with given ID
	//	GetMovie(int) (Movie, error)
	// GetAllMovies returns all movies saved in storage
	GetAllMovies() ([]Movie, error)
	// GetAllCastMembers returns a list of all cast members for a given movie ID
	//	GetAllCastMembers(int) []CastMember
}

// Service provides beer and review listing operations
type Service interface {
	//	GetMovie(int) (Movie, error)
	GetMovies() ([]Movie, error)
	//	GetMovieCastMembers(int) []CastMember
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
