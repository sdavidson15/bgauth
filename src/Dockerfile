FROM golang:1.20
WORKDIR /go/src/bgauth/
ADD . .
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/datastructures/main .

CMD ["./main"]
