This is a REST api, that allows for uploading and managing books in digital library.

## Running
### Database
There is a `docker-compose.yaml` config in directory `postgres`, which you can use.
Change the default login credentials if you want
### .env
Run `cp .env.example .env`, and edit the environemnt variables as needed
### Run!
Install [Go](https://go.dev/) and run `go run .`