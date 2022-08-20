# rental-service

The task is to develop a rentals JSON API that returns a list of rentals that can be filtered, sorted, and paginated. We have included files to create a database of rentals.

Your application should support the following endpoints.

- `/rentals/<RENTAL_ID>` Read one rental endpoint
- `/rentals` Read many (list) rentals endpoint
    - Supported query parameters
        - price_min (number)
        - price_max (number)
        - limit (number)
        - offset (number)
        - ids (comma separated list of rental ids)
        - near (comma separated pair [lat,lng])
        - sort (string)
    - Examples:
        - `rentals?price_min=9000&price_max=75000`
        - `rentals?limit=3&offset=6`
        - `rentals?ids=3,4,5`
        - `rentals?near=33.64,-117.93` // within 100 miles
        - `rentals?sort=price`
        - `rentals?near=33.64,-117.93&price_min=9000&price_max=75000&limit=3&offset=6&sort=price`

The rental object JSON in the response should have the following structure:
```json
{
  "id": "int",
  "name": "string",
  "description": "string",
  "type": "string",
  "make": "string",
  "model": "string",
  "year": "int",
  "length": "decimal",
  "sleeps": "int",
  "primary_image_url": "string",
  "price": {
    "day": "int"
  },
  "location": {
    "city": "string",
    "state": "string",
    "zip": "string",
    "country": "string",
    "lat": "decimal",
    "lng": "decimal"
  },
  "user": {
    "id": "int",
    "first_name": "string",
    "last_name": "string"
  }
}
```

## Run Application

```
$ docker-compose up --build -d
$ curl -i "http://localhost:3000/rentals"
```

## Execute Tests

```
$ go generate ./... 
$ go test -cover ./...
```
