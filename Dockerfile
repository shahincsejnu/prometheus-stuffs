FROM golang:1.15-alpine as go-builder

WORKDIR /go/src/github.com/shahincsejnu/prometheus-stuffs

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o prometheus-demo *.go


FROM alpine:latest

RUN mkdir -p /app && \
    addgroup -S app && adduser -S app -G app && \
    chown app:app /app

WORKDIR /app

COPY --from=go-builder /go/src/github.com/shahincsejnu/prometheus-stuffs .

USER app

CMD ["./prometheus-demo"]