-- Populates the database with some values

-- Create a couple of categories
INSERT INTO genres (name) VALUES ('Science Fiction');
INSERT INTO genres (name) VALUES ('Mystery');

-- Create a couple of books and assign them to categories
INSERT INTO books (name, author, pages, filename, genre_id)
VALUES ('1984', 'George Orwell', 353, '/var/www/books/1984.pdf', 1);
INSERT INTO books (name, author, pages, filename, genre_id)
VALUES ('The Girl with the Dragon Tattoo', 'Stieg Larsson', 465, '/var/www/books/girl_with_dragon_tattoo.pdf', 2);

-- Create a couple of shelves
INSERT INTO shelves (name) VALUES ('To Read');
INSERT INTO shelves (name) VALUES ('Favorites');

-- Assign the first book to the 'To Read' shelf
INSERT INTO book_shelf (book_id, shelf_id)
VALUES (1, 1);

-- Assign the second book to both 'To Read' and 'Favorites' shelves
INSERT INTO book_shelf (book_id, shelf_id)
VALUES (2, 1), (2, 2);
