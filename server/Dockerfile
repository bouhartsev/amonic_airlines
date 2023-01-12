FROM golang:1.19-alpine3.16

RUN mkdir -p /app/amonic

COPY . /app/amonic/server

WORKDIR /app/amonic/server

RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init

RUN go build -o server main.go

EXPOSE 3000

CMD ["./server"]
