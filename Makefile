schema:
	protoc --go_out=./go/grpc_demo/ movie_catalogue.proto
	protoc --python_out=./python/grpc_demo/ movie_catalogue.proto

clean:
	rm -f ./python/grpc_demo/*_pb2.py ./go/grpc_demo/*.pb.go
	
