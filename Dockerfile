FROM golang:1.23.2 AS builder

WORKDIR /app

COPY . ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /application

FROM scratch

COPY --from=builder /application /application

EXPOSE 8080

# Run
CMD ["/application"]