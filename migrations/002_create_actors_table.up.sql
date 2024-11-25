CREATE TABLE actors (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(100) NOT NULL,
                        gender CHAR(1) CHECK (gender IN ('M', 'F')),
                        birthdate DATE NOT NULL
);
