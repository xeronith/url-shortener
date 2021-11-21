FROM golang:1.17.3 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' -o ./bin/url-shortener .

FROM scratch
ENV BASE_URL="https://yourdomain.com"
ENV PORT=8080
COPY --from=builder /app/bin /app
EXPOSE $PORT
ENTRYPOINT ["/app/url-shortener"]