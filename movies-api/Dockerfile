# syntax=docker/dockerfile:1

##
## Build the application from source
##

FROM golang:1.19 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./


RUN CGO_ENABLED=0 GOOS=linux go build -o /movies-api

##
## Deploy the application binary into a lean image
##

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /movies-api /movies-api

USER nonroot:nonroot

