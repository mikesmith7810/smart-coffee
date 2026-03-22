FROM golang:1.25-alpine AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o smart-coffee .

FROM alpine:3.21
WORKDIR /app
COPY --from=builder /src/smart-coffee ./smart-coffee
EXPOSE 8080
ENTRYPOINT ["./smart-coffee"]
