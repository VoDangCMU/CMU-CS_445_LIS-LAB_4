FROM golang:1.23-alpine AS builder
WORKDIR /lab4
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /lab4
COPY --from=builder /lab4/main .
EXPOSE 8080
CMD ["./main"]
