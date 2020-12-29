package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jovv/grpc_demo/go/grpc_demo/http/rest"
)

func main() {

	lister := listing.NewService(env)

	router := rest.Handler(lister)

	fmt.Println("Distributing movies at: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
