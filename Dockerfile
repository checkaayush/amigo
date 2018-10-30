FROM golang:1.11 AS builder

# Download and install the latest release of dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/socialcopsdev/amigo
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app .

FROM scratch
COPY --from=builder /app ./
EXPOSE 5000
ENTRYPOINT ["./app"]