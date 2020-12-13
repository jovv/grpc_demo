# generate code for interoperability with binary format using the protocol buffer compiler 
schema:
	protoc --go_out=./go/grpc_demo/ movie_catalogue.proto
	protoc --python_out=./python/grpc_demo/ movie_catalogue.proto

# setup Python code with Poetry project (https://python-poetry.org/docs/)
setup:
	curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py | python -
	(cd python/grpc_demo && poetry install)

# build Go code
build:
	(cd go/grpc_demo && go build)

# cleanup generated code of protobuf compiler
clean:
	rm -f ./python/grpc_demo/*_pb2.py ./go/grpc_demo/*.pb.go
	
