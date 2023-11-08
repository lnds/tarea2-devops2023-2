CREATE TABLE movies(
    id serial PRIMARY KEY,
    title varchar(150) NOT NULL,
    year int NOT NULL,
    description text,
    director_id integer,
    constraint fk_movies_directors
        foreign key (director_id)
        REFERENCES directors(id)
)