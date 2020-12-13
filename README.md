# gRPC demo

Demo protobuf and gRPC concepts using a simple movie catalogue.

You can run the **PRESENT** markdown presentation with [patat](https://github.com/jaspervdj/patat).

## Movie Catalogue

Code generated from the protobuf schema in both Python and Go is already provided. Should you wish to do so, you can clean up the generated code with `make clean` or regenerate it with `make schema`.

An interface to add movies is provided in Python (for no particular reason).
Listing movies is possible with both Python and Go and demonstrates the language agnostic nature of the *protobuf* format.
The Python section uses version `3.9`.

## Setting things up

### Python

The Python section of this demo was created with [poetry](https://python-poetry.org/docs/). An automated setup is provided by simply running `make setup`.

From the `./python/grpc_demo/` folder, you can activate the virtual environment with `poetry shell`, which is required to run the Python interfaces.

### Go

You can build a binary with `make build`, it will assume the default binary name of `grpc_demo`.


## Interfacing with a Movie Catalogue

For the Python interfaces, verify you've activated the virtual environment with `poetry env info`.

### Adding

To add a movie, from `./python/grpc_demo` run

```bash
python add_movie.py <your catalogue file>
```

If the specified file does not yet exist, it will be created.

### Listing

To list the movies in the catalogue, from `./python/grpc_demo` run

```bash
python list_movie.py <your catalogue file>
```

or for the Go interface, from `./go/grpc_demo` run

```bash
./grpc_demo <your catalogue file>
```

For the Go interface, you should build the binary first with `make build` from the repo root. 

