schema:
	protoc --go_out=./go/grpc_demo/ movie_catalogue.proto
	protoc --python_out=./python/grpc_demo/ movie_catalogue.proto

setup:
	curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py | python -
	(cd python/grpc_demo && poetry install)

build:
	(cd go/grpc_demo && go build)

clean:
	rm -f ./python/grpc_demo/*_pb2.py ./go/grpc_demo/*.pb.go
	
