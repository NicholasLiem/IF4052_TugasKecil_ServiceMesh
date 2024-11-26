FROM golang:1.23 as builder

WORKDIR /app
COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add curl
COPY --from=builder /app/main /main
CMD ["/main"]
