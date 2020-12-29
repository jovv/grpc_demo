TRUNCATE TABLE movies;
TRUNCATE TABLE castmembers;

INSERT INTO movies 
(id, title, description, production_year, genre, duration) 
VALUES 
(1, 'Patser', 'Losjes gebaseerd op waargebeurde shit', 2018, 'Drama', 125),
(2, 'The Imitation Game', 
 'During World War II, the English mathematical genius Alan Turing tries to crack the German Enigma code with help from fellow mathematicians',
 2014, 'Historical drama', 114);

INSERT INTO castmembers 
(movie_id, character, first_name, last_name)
VALUES 
(1, 'Adamo', 'Matteo', 'Simoni'),
(2, 'Alan Turing', 'Benedict', 'Cumberbatch');