# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: movie_catalogue.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='movie_catalogue.proto',
  package='grpc',
  syntax='proto3',
  serialized_options=b'Z4github.com/jovv/grpc_demo/go/grpc_demo/pkg/http/grpc',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x15movie_catalogue.proto\x12\x04grpc\"U\n\nCastMember\x12\x11\n\tCharacter\x18\x01 \x01(\t\x12\x11\n\tGirstName\x18\x02 \x01(\t\x12\x10\n\x08LastName\x18\x03 \x01(\t\x12\x0f\n\x07MovieID\x18\x04 \x01(\x05\"\x97\x01\n\x05Movie\x12\n\n\x02ID\x18\x01 \x01(\x05\x12\r\n\x05Title\x18\x02 \x01(\t\x12\x13\n\x0b\x44\x65scription\x18\x03 \x01(\t\x12\x16\n\x0eProductionYear\x18\x04 \x01(\x05\x12\r\n\x05Genre\x18\x05 \x01(\t\x12\x10\n\x08\x44uration\x18\x06 \x01(\x05\x12%\n\x0b\x43\x61stMembers\x18\x07 \x03(\x0b\x32\x10.grpc.CastMember\"\x1a\n\x0cMovieRequest\x12\n\n\x02id\x18\x01 \x01(\x05\"(\n\nMovieReply\x12\x1a\n\x05movie\x18\x01 \x01(\x0b\x32\x0b.grpc.Movie\"\x0f\n\rMoviesRequest\"*\n\x0bMoviesReply\x12\x1b\n\x06movies\x18\x01 \x03(\x0b\x32\x0b.grpc.Movie2G\n\x0eMovieCatalogue\x12\x35\n\tGetMovies\x12\x13.grpc.MoviesRequest\x1a\x11.grpc.MoviesReply\"\x00\x42\x36Z4github.com/jovv/grpc_demo/go/grpc_demo/pkg/http/grpcb\x06proto3'
)




_CASTMEMBER = _descriptor.Descriptor(
  name='CastMember',
  full_name='grpc.CastMember',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='Character', full_name='grpc.CastMember.Character', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='GirstName', full_name='grpc.CastMember.GirstName', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='LastName', full_name='grpc.CastMember.LastName', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MovieID', full_name='grpc.CastMember.MovieID', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=31,
  serialized_end=116,
)


_MOVIE = _descriptor.Descriptor(
  name='Movie',
  full_name='grpc.Movie',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='ID', full_name='grpc.Movie.ID', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Title', full_name='grpc.Movie.Title', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Description', full_name='grpc.Movie.Description', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ProductionYear', full_name='grpc.Movie.ProductionYear', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Genre', full_name='grpc.Movie.Genre', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Duration', full_name='grpc.Movie.Duration', index=5,
      number=6, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='CastMembers', full_name='grpc.Movie.CastMembers', index=6,
      number=7, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=119,
  serialized_end=270,
)


_MOVIEREQUEST = _descriptor.Descriptor(
  name='MovieRequest',
  full_name='grpc.MovieRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='grpc.MovieRequest.id', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=272,
  serialized_end=298,
)


_MOVIEREPLY = _descriptor.Descriptor(
  name='MovieReply',
  full_name='grpc.MovieReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='movie', full_name='grpc.MovieReply.movie', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=300,
  serialized_end=340,
)


_MOVIESREQUEST = _descriptor.Descriptor(
  name='MoviesRequest',
  full_name='grpc.MoviesRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=342,
  serialized_end=357,
)


_MOVIESREPLY = _descriptor.Descriptor(
  name='MoviesReply',
  full_name='grpc.MoviesReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='movies', full_name='grpc.MoviesReply.movies', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=359,
  serialized_end=401,
)

_MOVIE.fields_by_name['CastMembers'].message_type = _CASTMEMBER
_MOVIEREPLY.fields_by_name['movie'].message_type = _MOVIE
_MOVIESREPLY.fields_by_name['movies'].message_type = _MOVIE
DESCRIPTOR.message_types_by_name['CastMember'] = _CASTMEMBER
DESCRIPTOR.message_types_by_name['Movie'] = _MOVIE
DESCRIPTOR.message_types_by_name['MovieRequest'] = _MOVIEREQUEST
DESCRIPTOR.message_types_by_name['MovieReply'] = _MOVIEREPLY
DESCRIPTOR.message_types_by_name['MoviesRequest'] = _MOVIESREQUEST
DESCRIPTOR.message_types_by_name['MoviesReply'] = _MOVIESREPLY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CastMember = _reflection.GeneratedProtocolMessageType('CastMember', (_message.Message,), {
  'DESCRIPTOR' : _CASTMEMBER,
  '__module__' : 'movie_catalogue_pb2'
  # @@protoc_insertion_point(class_scope:grpc.CastMember)
  })
_sym_db.RegisterMessage(CastMember)

Movie = _reflection.GeneratedProtocolMessageType('Movie', (_message.Message,), {
  'DESCRIPTOR' : _MOVIE,
  '__module__' : 'movie_catalogue_pb2'
  # @@protoc_insertion_point(class_scope:grpc.Movie)
  })
_sym_db.RegisterMessage(Movie)

MovieRequest = _reflection.GeneratedProtocolMessageType('MovieRequest', (_message.Message,), {
  'DESCRIPTOR' : _MOVIEREQUEST,
  '__module__' : 'movie_catalogue_pb2'
  # @@protoc_insertion_point(class_scope:grpc.MovieRequest)
  })
_sym_db.RegisterMessage(MovieRequest)

MovieReply = _reflection.GeneratedProtocolMessageType('MovieReply', (_message.Message,), {
  'DESCRIPTOR' : _MOVIEREPLY,
  '__module__' : 'movie_catalogue_pb2'
  # @@protoc_insertion_point(class_scope:grpc.MovieReply)
  })
_sym_db.RegisterMessage(MovieReply)

MoviesRequest = _reflection.GeneratedProtocolMessageType('MoviesRequest', (_message.Message,), {
  'DESCRIPTOR' : _MOVIESREQUEST,
  '__module__' : 'movie_catalogue_pb2'
  # @@protoc_insertion_point(class_scope:grpc.MoviesRequest)
  })
_sym_db.RegisterMessage(MoviesRequest)

MoviesReply = _reflection.GeneratedProtocolMessageType('MoviesReply', (_message.Message,), {
  'DESCRIPTOR' : _MOVIESREPLY,
  '__module__' : 'movie_catalogue_pb2'
  # @@protoc_insertion_point(class_scope:grpc.MoviesReply)
  })
_sym_db.RegisterMessage(MoviesReply)


DESCRIPTOR._options = None

_MOVIECATALOGUE = _descriptor.ServiceDescriptor(
  name='MovieCatalogue',
  full_name='grpc.MovieCatalogue',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=403,
  serialized_end=474,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetMovies',
    full_name='grpc.MovieCatalogue.GetMovies',
    index=0,
    containing_service=None,
    input_type=_MOVIESREQUEST,
    output_type=_MOVIESREPLY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_MOVIECATALOGUE)

DESCRIPTOR.services_by_name['MovieCatalogue'] = _MOVIECATALOGUE

# @@protoc_insertion_point(module_scope)
