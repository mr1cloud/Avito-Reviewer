FROM golang:latest AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /reviewer-service ./cmd/reviewer/main.go


FROM alpine:latest

COPY --from=builder /reviewer-service /home/reviewer-service

WORKDIR /home

ENTRYPOINT ["./reviewer-service"]

EXPOSE 8080