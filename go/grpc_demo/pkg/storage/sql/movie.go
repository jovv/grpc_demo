package sql

// Movie defines properties for the sql storage of cast member information
type Movie struct {
	ID             int
	Title          string
	Description    string
	ProductionYear int
	Genre          string
	Duration       int
}
