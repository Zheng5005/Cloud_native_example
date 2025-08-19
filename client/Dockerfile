# Build
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/client ./client

# Final
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /out/client /app/client
ENV TARGET_ADDR=grpc-server:50051
ENTRYPOINT ["/app/client"]
