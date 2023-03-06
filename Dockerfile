#build stage
FROM golang:1.20-alpine AS builder

ENV GO111MODULE=on \
#apline linux does not have gcc
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

#or use [RUN git clone ....]
COPY . .

RUN go mod download
RUN go build -a -o /app/todo-docker-test 

#deploy stage
FROM alpine:3.14

WORKDIR /app
COPY --from=builder /app/todo-docker-test /app/todo-docker-test

#APP ENV VARS (DB info)
ENV TODO_DB_USER=superuser \
    TODO_DB_PASS=superuser \
    TODO_DB_NAME=todo \
    TODO_DB_ADDRESS=192.168.0.33 \
    TODO_DB_PORT=3306 \
    TODO_DB_QUERY_MAX_LIMIT=500

EXPOSE 8083

CMD [ "/app/todo-docker-test" ]