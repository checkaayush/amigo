<p align="center">
  <img src="https://golangtraining-in-web.appspot.com/img/cowboy-color.png" height="130px"/>
  <br><br>
</p>

# amigo

A simple dockerized API starter for Golang.

Personal experimental playground to explore the Golang ecosystem.

## Development (with Docker and Docker Compose)

> Pre-requisites: Install latest stable versions of Docker and Docker Compose.

1. Clone the repository locally.
2. From repository root, run:
```bash
docker-compose up
```
3. API will be up and running at http://localhost:5000.

> Update all direct and indirect dependencies: `go get -u`
> Remove unused or no-longer needed dependencies: `go mod tidy`

## Built With

* [Echo](https://echo.labstack.com/) - The web framework used
* [viper](https://github.com/spf13/viper) - Configuration management

## Authors

* [Aayush Sarva](https://github.com/checkaayush)

## References

1. [Go Modules](https://github.com/golang/go/wiki/Modules)
