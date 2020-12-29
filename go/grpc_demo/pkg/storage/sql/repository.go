package sql

import (
	"database/sql"

	"github.com/jovv/grpc_demo/go/grpc_demo/pkg/listing"

	// postgres driver
	_ "github.com/lib/pq"
)

// Storage wraps the sql.DB connection pool
type Storage struct {
	db *sql.DB
}

// NewStorage returns a new sql storage
func NewStorage() (*Storage, error) {

	// TODO: move to some sort of config file, stup in main (read with cleanenv)
	db, err := sql.Open("postgres", "postgres://demo:demopw@localhost:35432/mc?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil

}

// GetAllMovies runs a query to retrieve all movies
func (s *Storage) GetAllMovies() ([]listing.Movie, error) {
	rows, err := s.db.Query("SELECT * FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := []listing.Movie{}

	for rows.Next() {
		var m Movie
		var movie listing.Movie

		err := rows.Scan(&m.ID, &m.Title, &m.Description, &m.ProductionYear, &m.Genre, &m.Duration)

		if err != nil {
			return nil, err
		}

		movie.ID = m.ID
		movie.Title = m.Title
		movie.Description = m.Description
		movie.ProductionYear = m.ProductionYear
		movie.Genre = m.Genre
		movie.Duration = m.Duration

		list = append(list, movie)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return list, nil

}
