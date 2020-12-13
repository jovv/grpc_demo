# gRPC demo

Demo protobuf and gRPC concepts using a simple movie catalogue.

Run the **PRESENT** markdown presentation with [patat](https://github.com/jaspervdj/patat).

# Python

## Poetry

The python section of this demo was created with [poetry](https://python-poetry.org/docs/).

You can find installation instructions [here](https://python-poetry.org/docs/#installation).

Then from `./python/grpc_demo/` run `poetry install` and activate the virtual environment with `poetry shell`.


## Movie Catalogue interface

To add a movie

```bash
python add_movie.py <your catalogue file>
```

If the specified file does not yet exist, it will be created.

To list the moves in the catalogue

```bash
python list_movie.py <your catalogue file>
```

# Go

Make sure you build the binary first. From the repo root you can 

```bash
make build
```

## Movie Catalogue interface

To list the moves in the catalogue

```bash
python list_movie.py <your catalogue file>
``` 

