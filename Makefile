# generate code for interoperability with binary format using the protocol buffer compiler 
GO_GRPC_PKG = github.com/jovv/grpc_demo/go/grpc_demo/pkg/http/grpc
GO_GRPC_PATH = go/grpc_demo/pkg/http/grpc
schema:
	protoc --go_opt=module=${GO_GRPC_PKG} \
		   --go_out=${GO_GRPC_PATH} \
		   --go-grpc_opt=module=${GO_GRPC_PKG} \
		   --go-grpc_out=${GO_GRPC_PATH} \
		   movie_catalogue.proto
	protoc --python_out=python/grpc_demo/ movie_catalogue.proto

# setup Python code with Poetry project (https://python-poetry.org/docs/)
setup:
	curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py | python -
	(cd python/grpc_demo && poetry install)

# build Go code
build:
	(cd go/grpc_demo/cmd/rest/movie_server && GOOS=linux go build)
	(cd go/grpc_demo/cmd/grpc/movie_server && GOOS=linux go build)

ddl:
	docker exec grpc_demo_postgres_1 psql -h localhost -U demo -d mc -f /scripts/movies_ddl.sql

dml:
	docker exec grpc_demo_postgres_1 psql -h localhost -U demo -d mc -f /scripts/movies_dml.sql

# cleanup generated code of protobuf compiler
clean:
	rm -f ./python/grpc_demo/*_pb2.py ./go/grpc_demo/movies/*.pb.go
