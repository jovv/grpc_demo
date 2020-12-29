package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jovv/grpc_demo/go/grpc_demo/pkg/http/rest"
	"github.com/jovv/grpc_demo/go/grpc_demo/pkg/listing"
	"github.com/jovv/grpc_demo/go/grpc_demo/pkg/storage/sql"
)

func main() {

	var lister listing.Service

	s, err := sql.NewStorage()
	if err != nil {
		log.Fatal(err)
	}
	lister = listing.NewService(s)

	router := rest.Handler(lister)

	fmt.Println("Showing movies at: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
