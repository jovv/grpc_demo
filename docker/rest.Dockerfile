FROM scratch

COPY movie_server /

EXPOSE 8080

ENTRYPOINT [ "/movie_server" ]