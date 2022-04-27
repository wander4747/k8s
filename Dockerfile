FROM golang:1.18
WORKDIR /go/src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server server.go
CMD ["./server"]