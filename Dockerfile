# Stage 1: Build
FROM golang:1.22-alpine3.18 AS builder

ARG APP_ENV=production

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Use build arg in build command if needed (example: conditional build)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# Stage 2: Run
FROM alpine:3.18

ARG TZ=Asia/Dhaka
ENV TZ=${TZ}

WORKDIR /root/

RUN apk add --no-cache tzdata ca-certificates && update-ca-certificates

COPY --from=builder /app/main .
COPY --from=builder /app/views ./views
COPY --from=builder /app/static ./static

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

EXPOSE 8080

CMD ["./main"]
