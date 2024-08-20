# Weather API

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

Sample weather API using data from the [National Weather Service](https://www.weather.gov/documentation/services-web-api).

## :rocket: Run

### Native

1. (Optional) Set port environment variable in your terminal or a `.env` file according to [.env.example](.env.example).
2. [Install Go](https://go.dev/doc/install).
3. Run unit tests:

```shell
go test ./...
```

4. Run with Go:

```shell
go run cmd/main.go
```

5. Test manually:

```shell
curl "http://localhost:8080/weather?lat=27.950575&lon=-82.457176"
```

### Docker

1. (Optional) Set port environment variable in [docker-compose.yaml](docker-compose.yaml) according to [.env.example](.env.example).
2. [Install Docker](https://docs.docker.com/engine/install).
3. Run with Docker compose:

```shell
docker-compose up
```

4. Test manually:

```shell
curl "http://localhost:8080/weather?lat=27.950575&lon=-82.457176"
```
