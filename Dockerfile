ARG GO_VERSION=1.23
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN make build


FROM debian:bookworm

COPY --from=builder /app/main /usr/local/bin/
CMD ["main"]
