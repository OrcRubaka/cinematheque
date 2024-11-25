CREATE TABLE movies (
                        id SERIAL PRIMARY KEY,
                        title VARCHAR(150) NOT NULL,
                        description TEXT CHECK (char_length(description) >= 1000),
                        release_date DATE NOT NULL,
                        rating NUMERIC(2, 1) CHECK (rating BETWEEN 0 AND 10)
);
