<div align="center">
  <p>
    <img src="https://golangtraining-in-web.appspot.com/img/cowboy-color.png" height="130px"/>
    <img>
  </p>
  <h1>amigo</h1>
  <a href="https://goreportcard.com/report/github.com/checkaayush/amigo">
    <img src="https://goreportcard.com/badge/github.com/checkaayush/amigo"/>
  </a>
</div>

## Introduction

A simple dockerized API starter for Golang inspired by [GORSK](https://github.com/ribice/gorsk).

Shout-out to [ribice](https://github.com/ribice) for the awesome work on GORSK v2. :thumbsup:

### What's different from GORSK?

* Uses docker and docker compose
* Has live code reloading for faster development
* Uses Go modules with Go 1.12+ and semantic versioning
* Uses Labstack Echo v4
* Uses MongoDB as the primary datastore along with the official MongoDB Go driver
* Uses environment variables for configuration
* Minor changes in repository structure

## Development

> Pre-requisites: Install latest stable versions of Docker and Docker Compose.

1. Clone the repository locally.
2. Add .env file in the repository root by modifying the .env.template file as needed.
3. From repository root, run:
```bash
docker-compose up
```
4. API will be up and running at http://localhost:5000.

### Dependency Management

Amigo uses Go modules with semantic versioning and is tested with Go 1.12+.

* Update all direct and indirect dependencies using `go get -u`.
* Remove unused dependencies using `go mod tidy`.
* Add a new dependency using `go get <path-to-dependency>`.

#### Dependencies

* [echo](https://echo.labstack.com/) - Web framework
* [envconfig](https://github.com/kelseyhightower/envconfig) - Configuration management
* [realize](https://github.com/oxequa/realize) - Live reloading
* [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) - MongoDB driver
* [testify](https://github.com/stretchr/testify) - Assertions library

## References

* [Go Modules](https://github.com/golang/go/wiki/Modules)
* [Using MongoDB Go Driver](https://vkt.sh/go-mongodb-driver-cookbook/)
