CREATE TABLE actors (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(100) NOT NULL,
                        gender CHAR(1) CHECK (gender IN ('M', 'F')),
                        birthdate DATE NOT NULL
);


CREATE TABLE movies (
                        id SERIAL PRIMARY KEY,
                        title VARCHAR(150) NOT NULL,
                        description TEXT CHECK (char_length(description) >= 1000),
                        release_date DATE NOT NULL,
                        rating NUMERIC(2, 1) CHECK (rating BETWEEN 0 AND 10)
);


CREATE TABLE movie_actors (
                              movie_id INT REFERENCES movies(id) ON DELETE CASCADE,
                              actor_id INT REFERENCES actors(id) ON DELETE CASCADE,
                              PRIMARY KEY (movie_id, actor_id)
);
