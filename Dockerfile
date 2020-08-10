FROM golang:1.14-alpine3.12 as builder

# todo build valid version for OS

WORKDIR /app
COPY . .
RUN go build -o /app main.go
RUN ls

FROM wordpress:cli-php7.4

WORKDIR /app
COPY wp-cli.yaml .

ENV WP_CLI_CONFIG_PATH=/app/wp-cli.yaml
ENTRYPOINT ["./main"]
COPY --from=builder /app/main .
