package grpc

import (
	context "context"

	"github.com/jovv/grpc_demo/go/grpc_demo/pkg/listing"
)

// Handler implements the proto MovieCatalogue service
type Handler struct {
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
func (mc *Handler) GetMovies(context.Context, *MoviesRequest) (*MoviesReply, error) {
	l, err := mc.Lister.GetMovies()
	if err != nil {
		return nil, err
	}

	return &MoviesReply{
		Movies: moviesAdapter(l),
	}, nil
}
