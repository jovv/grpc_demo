#! /usr/bin/env python
"""Iterate through all movies and print info about them"""

import sys
import movie_catalogue_pb2


def list_movies(movie_catalogue):
    """Iterate through all movies and print info about them"""
    for movie in movie_catalogue.movies:
        print("Title:", movie.title)
        print("Description:", movie.description)
        print("Production Year:", movie.productionYear)
        print("Genre:", movie.genre)
        print("Duration:", movie.duration)

        for cast_member in movie.castMembers:
            print("Character:", cast_member.character)
            print("First Name:", cast_member.firstName)
            print("Last Name:", cast_member.lastName)


def main():
    """Read the entire movie catalogue from a file and print all the
       information inside"""
    if len(sys.argv) != 2:
        print("Usage:", sys.argv[0], "MOVIE_CATALOGUE_FILE")
        sys.exit(-1)

    movie_catalogue = movie_catalogue_pb2.MovieCatalogue()

    # Read the existing movie catalogue
    with open(sys.argv[1], "rb") as f:
        movie_catalogue.ParseFromString(f.read())

    list_movies(movie_catalogue)
