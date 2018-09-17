# Build env
FROM golang:1.10.3 AS build

WORKDIR $GOPATH/src/github.com/arthurc0102/ntub-class-table-api
COPY . .

RUN go get -v -u github.com/golang/dep/cmd/dep
RUN dep ensure

# Ref: https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main.out

RUN cp ./main.out /

# Run env
FROM alpine:3.8

WORKDIR /app

COPY --from=build /main.out .
EXPOSE 8080

# COPY ./docker-entrypoint.sh /usr/local/bin/
# RUN chmod +x /usr/local/bin/docker-entrypoint.sh
# ENTRYPOINT [ "docker-entrypoint.sh" ]

CMD [ "./main.out" ]
