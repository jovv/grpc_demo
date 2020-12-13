package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
)

func writeMovie(w io.Writer, m *Movie) {
	fmt.Fprintln(w, "Title:", m.Title)
	fmt.Fprintln(w, "  Description:", m.Description)
	fmt.Fprintln(w, "  Production Year:", m.ProductionYear)
	fmt.Fprintln(w, "  Genre:", m.Genre)
	fmt.Fprintln(w, "  Duration:", m.Duration)

	for _, cm := range m.CastMembers {
		fmt.Fprintln(w, cm.Character)
		fmt.Fprintln(w, cm.FirstName)
		fmt.Fprintln(w, cm.LastName)
	}
}

func listMovies(w io.Writer, movieCatalogue *MovieCatalogue) {
	for _, m := range movieCatalogue.Movies {
		writeMovie(w, m)
	}
}

// Main reads the entire address book from a file and prints all the
// information inside.
func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage:  %s MOVIE_CATALOGUE_FILE\n", os.Args[0])
	}
	fname := os.Args[1]

	// [START unmarshal_proto]
	// Read the existing address book.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	movieCatalogue := &MovieCatalogue{}
	if err := proto.Unmarshal(in, movieCatalogue); err != nil {
		log.Fatalln("Failed to parse movie catalogue:", err)
	}
	// [END unmarshal_proto]

	listMovies(os.Stdout, movieCatalogue)
}
