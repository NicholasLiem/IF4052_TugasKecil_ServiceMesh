FROM golang:1.23 as builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM gcr.io/distroless/base

COPY --from=builder /app/main /

CMD ["/main"]
