# Build
FROM golang:1.22.6-bookworm AS builder
WORKDIR /app
COPY . .
RUN go get
RUN go build -o main

# Run
FROM debian:12.4
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .

EXPOSE 8080

CMD [ "/app/main" ]
