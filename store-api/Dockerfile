### STAGE 1: Build ###
FROM golang:alpine AS builder
WORKDIR /build
ADD ./store-api/go.mod ./
COPY ./store-api ./
RUN go build -o server ./cmd/main.go


### STAGE 2: Run ###
FROM alpine:3.18.2
COPY --from=builder /build/server server
COPY --from=builder /build/config.toml config.toml
RUN mkdir assets
CMD ["./server"]
