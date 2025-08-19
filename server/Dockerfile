# Build
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/server ./server

# Final
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /out/server /app/server
ENV LISTEN_ADDR=:50051
EXPOSE 50051
USER nonroot:nonroot
ENTRYPOINT ["/app/server"]
