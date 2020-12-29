package listing

// Repository provides access to the movie and castmembers storage
type Repository interface {
	// GetMovie returns the movie with given ID
	GetMovie(int) (Movie, error)
	// GetAllMovies returns all movies saved in storage
	GetAllMovies() []Movie
	// GetAllCastMembers returns a list of all cast members for a given movie ID
	GetAllCastMembers(int) []CastMember
}

// Service provides beer and review listing operations
type Service interface {
	GetMovie(int) (Movie, error)
	GetMovies() []Movie
	GetMovieCastMembers(int) []CastMember
}

type service struct {
	r Repository
}

// GetMovies returns all movies
func (s *service) GetMovies() []Movie {
	return s.r.GetAllMovies()
}
