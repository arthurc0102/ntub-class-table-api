# Build env
FROM golang:1.18.4 AS build

WORKDIR /srv

COPY . .

RUN go mod download

# Ref: https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# Run env
FROM alpine:3

COPY --from=build /srv/server /

ENV GIN_MODE=release

EXPOSE 8080

USER daemon

CMD [ "/server" ]
