# syntax=docker/dockerfile:1-labs
# -----------------------------------------------------
# Build
# -----------------------------------------------------
FROM golang:1.18-bullseye as base

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /workspace

COPY go.mod go.sum ./
RUN <<EOF
    go mod download
EOF

COPY . .
RUN <<EOF
    go build
EOF

# -----------------------------------------------------
# Run
# -----------------------------------------------------
FROM debian:stable-slim

WORKDIR /app
COPY --from=base /workspace/try-golang try-golang

CMD ["/app/try-golang", "-onetime", "-example", ""]
