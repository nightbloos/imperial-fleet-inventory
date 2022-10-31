FROM golang:1.17.1-buster as build

WORKDIR /src/services/http
COPY . /src
RUN go mod download
RUN	CGO_ENABLED=0 go build -o /bin/srv ./cmd/

FROM alpine:3 as production

COPY --from=build /bin/srv /app/srv
ENTRYPOINT /app/srv
