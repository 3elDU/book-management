-- Populates the database with some values

-- Create a couple of categories
INSERT INTO genres (name) VALUES ('Science Fiction');
INSERT INTO genres (name) VALUES ('Mystery');
INSERT INTO genres (name) VALUES ('IT');

-- Create a couple of books and assign them to categories
INSERT INTO books (name, author, pages, filename, genre_id)
VALUES ('1984', 'George Orwell', 353, '/var/www/books/1984.pdf', 1);
INSERT INTO books (name, author, pages, filename, genre_id)
VALUES ('The Girl with the Dragon Tattoo', 'Stieg Larsson', 465, '/var/www/books/girl_with_dragon_tattoo.pdf', 2);
INSERT INTO books (name, author, pages, filename, genre_id)
VALUES ('The C++ Programming Language, 4th Edition', 'Bjarne Stroustrup', 636, '/var/www/books/cpp.pdf', 3);
INSERT INTO books (name, author, pages, filename, genre_id)
VALUES ('Linux Bible', 'Cristopher Negus', 837, '/var/www/books/linux_bible.pdf', 3);

-- Create a couple of shelves
INSERT INTO shelves (name) VALUES ('To Read');
INSERT INTO shelves (name) VALUES ('Favorites');

-- Assign the first book to the 'To Read' shelf
INSERT INTO book_shelf (book_id, shelf_id)
VALUES (1, 1);

-- Assign the second book to both 'To Read' and 'Favorites' shelves
INSERT INTO book_shelf (book_id, shelf_id)
VALUES (2, 1), (2, 2);

INSERT INTO book_shelf (book_id, shelf_id)
VALUES (3, 1), (4, 1);