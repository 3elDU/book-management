CREATE TABLE genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    pages INTEGER NOT NULL,
    filename VARCHAR(255),
    genre_id INTEGER REFERENCES genres(id)
);

CREATE TABLE shelves (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE book_shelf (
    book_id INTEGER REFERENCES books(id),
    shelf_id INTEGER REFERENCES shelves(id),
    PRIMARY KEY (book_id, shelf_id)
);