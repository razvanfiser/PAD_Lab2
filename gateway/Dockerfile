FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main ./gateway.go

EXPOSE 5003

CMD ["./main"]