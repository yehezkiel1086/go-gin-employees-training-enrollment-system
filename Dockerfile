FROM golang:1.24 AS builder

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /go/bin/app cmd/main.go

FROM golang:1.24-alpine

WORKDIR /app

COPY --from=builder /go/bin/app .

# Copy the .env file. It's recommended to use a .env.docker or pass variables at runtime for production.
COPY .env .

EXPOSE 8080
ENV PORT=8080

CMD ["./app"]
