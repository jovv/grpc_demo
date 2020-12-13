---
title: Protocol Buffers and gRPC
author: Jochen Van de Voorde
purpose: tech demo for aviDci squad @ DPG Media

---

# Protocol Buffers

* language neutral data structure or schema, `.proto`
* serializes to a binary message format
* ...
* allows for comments

[docs](https://developers.google.com/protocol-buffers/docs/overview)


# Code generation

The [protocol buffer compiler](https://grpc.io/docs/protoc-installation/) generates code for your language of choice to (de-) serialize from and to the binary format.

`movie_catalogue.pb.go`

`movie_catalogue_pb2.py`

# Schema evolution

Rules for backward and forward compatibility

* you *must* not change the tag numbers of any existing fields
* you *may* delete fields
* you *may* add new fields but you must use fresh tag numbers (i.e. tag numbers that were never used in this protocol buffer, not even by deleted fields)

# Caveats

* not designed to handle large messages, RoT `< 1 MB`
* not a self-describing format, schema is separated from the content

# gRPC

TODO
