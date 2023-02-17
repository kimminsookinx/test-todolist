FROM golang:1.20-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

ADD . /app/

RUN go mod download

RUN go build -o /app/todo-docker-test

EXPOSE 8083

CMD [ "/app/todo-docker-test" ]