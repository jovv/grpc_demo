package sql

import (
	"database/sql"

	"github.com/jovv/grpc_demo/go/grpc_demo/pkg/listing"
)

// Storage wraps the sql.DB connection pool
type Storage struct {
	db *sql.DB
}

// NewStorage returns a new sql storage
func NewStorage() (*Storage, error) {

	db, err := sql.Open("postgres", "demo@db:5432/mc")
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

	// Execute the SQL query by calling the All() method.
	bks, err := s.GetMovies()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}