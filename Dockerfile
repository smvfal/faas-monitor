FROM golang:1.15-alpine AS builder
WORKDIR /go/src/github.com/smvfal/faas-monitor
COPY . .
RUN go build -o faas-monitor

FROM alpine
WORKDIR /root/
COPY --from=builder /go/src/github.com/smvfal/faas-monitor/faas-monitor .
CMD ["./faas-monitor"]