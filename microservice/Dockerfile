FROM golang:1.20-alpine as builder
WORKDIR /go/src/github.com/thabetj/devops-journey/
COPY go.mod go.sum main.go ./
RUN go build -o todo

FROM alpine:latest
RUN apk --no-cache add ca-certificates gcompat
WORKDIR /root/
COPY --from=builder /go/src/github.com/thabetj/devops-journey/todo .
CMD ["./todo"]
