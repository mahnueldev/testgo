# Use a Go base image
FROM golang:1.22


WORKDIR /testgo
ENV CGO_ENABLED=1


COPY go.mod go.sum ./


RUN go install github.com/cosmtrek/air@latest

COPY . ./
CMD ["air", "-c", ".air.toml"]


