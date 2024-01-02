# Stage 1: Build the Go application
FROM golang:1.20.4-alpine as builder

WORKDIR /app

RUN apk add --no-cache build-base git



COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

# Stage 2: Create the final image using Alpine
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/


COPY --from=builder /app/main .

CMD ["./main"]
