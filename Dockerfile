FROM golang:1.15-alpine

WORKDIR /go/src/github.com/hanami125/todo-app
COPY . .

CMD ["go", "run", "main.go"]
