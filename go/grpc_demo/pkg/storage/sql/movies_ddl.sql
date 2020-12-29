CREATE TABLE movies (
    id int,
    title text,
    description text,
    production_year int,
    genre text,
    duration int    
);

CREATE TABLE castmembers (
    movie_id int,
    character text,
    first_name text,
    last_name text
);