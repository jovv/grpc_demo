package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	rpc "github.com/jovv/grpc_demo/go/grpc_demo/pkg/http/grpc"

	"github.com/jovv/grpc_demo/go/grpc_demo/pkg/listing"
	"github.com/jovv/grpc_demo/go/grpc_demo/pkg/storage/sql"
)

func main() {
	s, err := sql.NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	lister := listing.NewService(s)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := &rpc.Server{
		Lister: lister,
	}

	srv := grpc.NewServer()
	rpc.RegisterMovieCatalogueServer(srv, server)
	// register reflection service on gRPC server, so we can list available
	// services etc
	reflection.Register(srv)
	fmt.Println("Producing movies at: http://localhost:50051")
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
