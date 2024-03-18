FROM golang:latest
LABEL authors="almas"

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o film_catalog_app ./cmd/main.go

EXPOSE 8080

CMD ["./film_catalog_app"]