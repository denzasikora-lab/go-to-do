# syntax=docker/dockerfile:1
# Multi-stage image: compile a static Linux binary, then ship a minimal runtime.
FROM golang:1.22-alpine AS build
RUN apk add --no-cache git ca-certificates
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Static binary for distroless-like deployment (no libc dependency in final image).
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/todobot ./cmd/bot

FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata
COPY --from=build /out/todobot /usr/local/bin/todobot
USER nobody
ENTRYPOINT ["/usr/local/bin/todobot"]
