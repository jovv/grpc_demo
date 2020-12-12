#! /usr/bin/env python

import movie_catalogue_pb2
import sys

# Iterates through all movies and prints info about them
def list_movies(movie_catalogue):
  for movie in movie_catalogue.movies:
    print("Title:", movie.title)
    print("Description:", movie.description)
    print("Production Year:", movie.productionYear) 

    for cast_member in movie.castMembers:
        print("Character:", cast_member.character)
        print("First Name:", cast_member.firstName)
        print("Last Name:", cast_member.lastName)


    print("Genre:", movie.genre)
    print("Duration:", movie.duration)

def main():

    # Read the entire movie catalogue from a file and print all the 
    # information inside
    if len(sys.argv) != 2:
      print("Usage:", sys.argv[0], "MOVIE_CATALOGUE_FILE")
      sys.exit(-1)

    movie_catalogue = movie_catalogue_pb2.MovieCatalogue()

    # Read the existing address book.
    with open(sys.argv[1], "rb") as f:
    movie_catalogue.ParseFromString(f.read())

    list_movies(movie_catalogue)

