## How to run
First, you need postgresql server.
You can either install and configure it manually, or use a provided `docker-compose.yaml` file.
For the second case, navigate to directory `postgres`, and run  
`docker compose up -d`  
Now, go to project root, and simply run `go run .`