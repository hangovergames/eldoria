# Build Stage
FROM golang:1.21.6 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o eldoria ./cmd/eldoria

# Final Stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/eldoria .
COPY ./mocks/mock-systemctl.sh /usr/bin/systemctl
RUN chmod +x /usr/bin/systemctl
EXPOSE 9500
CMD ["./eldoria"]
