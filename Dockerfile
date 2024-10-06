FROM golang:alpine AS build

ENV GOOS linux

WORKDIR /build

ADD go.mod .
COPY . .

RUN go build -o out/main -v cmd/quotes/main.go

FROM alpine

WORKDIR /api

COPY --from=build /build/out/main /api/main

ENTRYPOINT ["./main"]