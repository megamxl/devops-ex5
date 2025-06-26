FROM golang:1.24-bookworm AS builder

WORKDIR /build

COPY .. .

ENV CGO_ENABLED 0
ENV GOOS linux

RUN go mod tidy

RUN go build -o serverBinary

FROM alpine:3.20.6

WORKDIR /app

COPY --from=builder /build/serverBinary /app/serverBinary

EXPOSE 8080

CMD ["/app/serverBinary"]