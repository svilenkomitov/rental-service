# rental-service

# Run

```
$ docker-compose up --build -d

$ curl -i "http://localhost:3000/rentals/1"
```

# Run Tests

```
$ go generate ./... 
$ go test ./...
```
