FROM golang:1.12.4 AS builder

RUN go get github.com/oxequa/realize
ENV GO111MODULE=on

# Download dependencies
WORKDIR /usr/src/app
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code from the host and build it
COPY . /usr/src/app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app ./cmd/api/main.go

EXPOSE 5000
CMD [ "realize", "start" ]
