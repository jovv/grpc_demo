package grpc

import (
	context "context"

	"github.com/jovv/grpc_demo/go/grpc_demo/pkg/listing"
)

// Server implements the proto MovieCatalogue service
type Server struct {
	Lister listing.Service
	UnimplementedMovieCatalogueServer
}

func movieAdapter(m listing.Movie) *Movie {
	return &Movie{
		ID:             int32(m.ID),
		Title:          m.Title,
		Description:    m.Description,
		ProductionYear: int32(m.ProductionYear),
		Genre:          m.Genre,
		Duration:       int32(m.Duration),
	}
}

func moviesAdapter(ms []listing.Movie) []*Movie {
	movies := make([]*Movie, len(ms))
	for i, m := range ms {
		movies[i] = movieAdapter(m)
	}
	return movies
}

// GetMovies returns an RPC MoviesReply with a list of all movies
func (mc *Server) GetMovies(context.Context, *MoviesRequest) (*MoviesReply, error) {
	l, err := mc.Lister.GetMovies()
	if err != nil {
		return nil, err
	}

	return &MoviesReply{
		Movies: moviesAdapter(l),
	}, nil
}

// GetMovie returns an RPC MovieReply with a movie for a specified ID
func (mc *Server) GetMovie(c context.Context, mr *MovieRequest) (*MovieReply, error) {
	m, err := mc.Lister.GetMovie(int(mr.ID))
	if err != nil {
		return nil, err
	}

	return &MovieReply{
		Movie: movieAdapter(m),
	}, nil
}
