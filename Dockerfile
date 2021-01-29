FROM golang:1.15-alpine as builder
ADD . /go/src/github.com/onuryilmaz/k8s-scheduler-extender
WORKDIR /go/src/github.com/onuryilmaz/k8s-scheduler-extender/cmd
RUN go build -v 
 
FROM alpine:latest
COPY --from=builder /go/src/github.com/onuryilmaz/k8s-scheduler-extender/cmd/cmd /usr/local/bin/k8s-scheduler-extender
CMD ["k8s-scheduler-extender"]