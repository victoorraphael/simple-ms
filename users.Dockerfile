FROM golang:alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY users/ .
COPY adapters/ .

RUN go build -o ./bin/users

ENTRYPOINT ["/app/bin/users"]
