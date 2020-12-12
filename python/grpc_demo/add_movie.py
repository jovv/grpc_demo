#! /usr/bin/env python
"""Add a Movie based on user input"""

import sys
import movie_catalogue_pb2


def prompt_for_movie(movie):
    """Fill in a Movie message based on user input"""

    movie.title = input("Enter title: ")

    description = input("Enter description (blank for none): ")
    if description != "":
        movie.description = description

    movie.productionYear = int(input("Enter production year: "))
    movie.genre = input("Enter genre: ")

    while True:
        print("Enter a cast member:")

        character = input("Enter a character or leave blank to finish): ")
        if character == "":
            break

        cast_member = movie.castMembers.add()
        cast_member.character = character

        first_name = input("Enter a first name: ")
        cast_member.firstName = first_name

        last_name = input("Enter a last name: ")
        cast_member.lastName = last_name


#  Read the entire movie catalogue from a file,
#  add one person based on user input, then write it back out to the same file
if len(sys.argv) != 2:
    print("Usage:", sys.argv[0], "MOVIE_CATALOGUE_FILE")
    sys.exit(-1)

movie_catalogue = movie_catalogue_pb2.MovieCatalogue()

# Read the existing movie catalogue
try:
    with open(sys.argv[1], "rb") as f:
        movie_catalogue.ParseFromString(f.read())
except IOError:
    print(sys.argv[1] + ": File not found.  Creating a new file.")

# Add a movie
prompt_for_movie(movie_catalogue.movies.add())

# Write the new movie catalogue back to disk
with open(sys.argv[1], "wb") as f:
    f.write(movie_catalogue.SerializeToString())
