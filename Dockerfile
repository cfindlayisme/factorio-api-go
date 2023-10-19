FROM golang:1.20.10 AS builder

WORKDIR /app

COPY * ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /application

FROM alpine:3.18

COPY --from=builder /application /application

EXPOSE 8080

# Run
CMD ["/application"]