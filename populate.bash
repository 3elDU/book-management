#! /usr/bin/env sh

curl -H "Content-Type: application/json" -X POST -d '{"name":"To Kill a Mockingbird","author":"Harper Lee","pages":283}' http://localhost:8080/new_book
curl -H "Content-Type: application/json" -X POST -d '{"name":"1984","author":"George Orwell","pages":328}' http://localhost:8080/new_book
curl -H "Content-Type: application/json" -X POST -d '{"name":"The Great Gatsby","author":"F. Scott Fitzgerald","pages":180}' http://localhost:8080/new_book
curl -H "Content-Type: application/json" -X POST -d '{"name":"Pride and Prejudice","author":"Jane Austen","pages":279}' http://localhost:8080/new_book
curl -H "Content-Type: application/json" -X POST -d '{"name":"The Catcher in the Rye","author":"J.D. Salinger","pages":224}' http://localhost:8080/new_book
curl -H "Content-Type: application/json" -X POST -d '{"name":"Brave New World","author":"Aldous Huxley","pages":288}' http://localhost:8080/new_book
curl -H "Content-Type: application/json" -X POST -d '{"name":"Lord of the Flies","author":"William Golding","pages":224}' http://localhost:8080/new_book
curl -H "Content-Type: application/json" -X POST -d '{"name":"One Hundred Years of Solitude","author":"Gabriel Garcia Marquez","pages":417}' http://localhost:8080/new_book
curl -H "Content-Type: application/json" -X POST -d '{"name":"Animal Farm","author":"George Orwell","pages":112}' http://localhost:8080/new_book
curl -H "Content-Type: application/json" -X POST -d '{"name":"The Hobbit","author":"J.R.R. Tolkien","pages":320}' http://localhost:8080/new_book
