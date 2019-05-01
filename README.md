<p align="center">
  <img src="https://golangtraining-in-web.appspot.com/img/cowboy-color.png" height="130px"/>
  <br><br>
</p>

# amigo

A simple dockerized API starter for Golang.

Personal experimental playground to explore the Golang ecosystem.

## Development

> Pre-requisites: Install latest stable versions of Docker and Docker Compose.

1. Clone the repository locally.
2. Add .env file in the repository root by modifying the .env.template file as needed.
3. From repository root, run:
```bash
docker-compose up
```
4. API will be up and running at http://localhost:5000.

> Update all direct and indirect dependencies: `go get -u`

> Remove unused dependencies: `go mod tidy`

## Built With

* [Echo](https://echo.labstack.com/) - Web framework
* [Envconfig](https://github.com/kelseyhightower/envconfig) - Configuration management
* [Realize](https://github.com/oxequa/realize) - Live reloading

## Authors

* [Aayush Sarva](https://github.com/checkaayush)

## References

* [Go Modules](https://github.com/golang/go/wiki/Modules)
