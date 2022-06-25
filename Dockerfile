# syntax=docker/dockerfile:1
FROM golang:1.18-bullseye
RUN go install github.com/beego/bee/v2@latest

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_HOME /go/src/authapp
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"
# COPY go.mod WORKDIR
# COPY go.sum WORKDIR

# RUN go mod download
# COPY ./ ./
# RUN go build -o /product-go-micro
EXPOSE 9090
CMD ["bee", "run"]