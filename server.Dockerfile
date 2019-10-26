FROM golang:latest as builder
WORKDIR /app

COPY ./server ./server
COPY go.mod .
COPY go.sum .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./server

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
