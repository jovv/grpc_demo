package listing

// Movie defines properties of a movie that should be included in listing movie
// information
type Movie struct {
	ID             int
	Title          string
	Description    string
	ProductionYear int
	Genre          string
	Duration       int
}
