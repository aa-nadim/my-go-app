# my-go-app

## Setup
```
go mod init my-go-app
go mod tidy
go run main.go

```

## Run Container

```
docker compose up --build -d
```


## Docker Hub
```
docker login
docker build -t a2nadim/go-mvc-app:latest .

docker push a2nadim/go-mvc-app:latest

docker tag a2nadim/go-mvc-app:latest a2nadim/go-mvc-app:v1.0
docker push a2nadim/go-mvc-app:v1.0

```

## Docker Image Layer Analysis
- Image: `a2nadim/go-mvc-app:latest`
- Total Layers: 12
- Final Image Size: ~25.7 MB
- Build Type: Multi-stage with Alpine base

### Layer Breakdown
| Layer   | Description                                                     | Size   |
| ------- | --------------------------------------------------------------- | ------ |
| 1       | CMD `["./main"]`                                                | 0B     |
| 2       | EXPOSE 8080                                                     | 0B     |
| 3       | `chown -R appuser:appgroup /root/main /root/views /root/static` | 8.37MB |
| 4       | `addgroup -S appgroup && adduser -S appuser`                    | 4.71kB |
| 5       | `COPY` binary from builder                                      | 1.43MB |
| 6       | `COPY` views (HTML)                                             | 16.4kB |
| 7       | `COPY` static assets (CSS/JS)                                   | 6.92MB |
| 8       | Install runtime deps (`tzdata`, `ca-certificates`)              | 1.62MB |
| 9       | Set `WORKDIR /root`                                             | 0B     |
| 10      | Set `ENV TZ=` (from `.env`)                                     | 0B     |
| 11      | Alpine base image                                               | 0B     |
| 12      | ADD `alpine-minirootfs`                                         | 7.36MB |


