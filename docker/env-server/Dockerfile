# syntax=docker/dockerfile:1
# https://docs.docker.com/language/golang/build-images/
FROM golang:1.16-alpine

# https://github.com/alpinelinux/docker-alpine/issues/98#issuecomment-679306519
RUN sed -i 's/https/http/' /etc/apk/repositories
RUN apk --no-cache add curl

WORKDIR /app

COPY main.go ./

EXPOSE 8080

CMD ["go","run","main.go"]