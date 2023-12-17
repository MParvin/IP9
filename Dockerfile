FROM golang:1.18.0-alpine3.14 AS builder

WORKDIR /app

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o ip9 main.go

FROM alpine:3.14

COPY --from=builder /app/ip9 /usr/local/bin/ip9

ENTRYPOINT ["/usr/local/bin/ip9"]
