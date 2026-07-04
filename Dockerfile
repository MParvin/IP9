FROM golang:1.20-alpine AS builder

WORKDIR /app

RUN apk add --no-cache curl

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN sh scripts/download-data.sh

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o ip9 main.go

FROM alpine:3.14

ARG PORT=8080

COPY --from=builder /app/ip9 /usr/local/bin/ip9
COPY --from=builder /app/data/*.xdb /data/

ENV IP2REGION_V4_DB=/data/ip2region_v4.xdb
ENV IP2REGION_V6_DB=/data/ip2region_v6.xdb

ENTRYPOINT ["/usr/local/bin/ip9"]
