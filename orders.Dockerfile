FROM golang:alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY orders/ .
COPY adapters/ .

RUN go build -o ./bin/orders

ENTRYPOINT ["/app/bin/orders"]
