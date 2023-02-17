FROM golang:1.20-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

#APP ENV VARS (DB info)
ENV TODO_DB_USER=superuser \
    TODO_DB_PASS=superuser \
    TODO_DB_NAME=todo \
    TODO_DB_ADDRESS=192.168.0.33 \
    TODO_DB_PORT=3306 \
    TODO_DB_QUERY_MAX_LIMIT=500

WORKDIR /app

ADD . /app/

RUN go mod download

RUN go build -o /app/todo-docker-test

EXPOSE 8083

CMD [ "/app/todo-docker-test" ]