FROM golang:1.22-alpine3.19
WORKDIR /usr/src/app

COPY ./ ./

RUN go mod download && go mod verify

RUN go build -o film_catalog_app ./cmd/main.go

CMD ["./film_catalog_app"]