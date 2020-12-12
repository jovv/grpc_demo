#	protoc -I=$(pwd) --go_out=$(pwd)/go/grpc_demo/ movie_catalogue.proto

#	protoc -I=$(pwd) --python_out=$(pwd)/python/grpc_demo/ movie_catalogue.proto

clean:
	rm -f ./python/grpc_demo/*_pb2.py ./go/grpc_demo/*.pb.go
	
