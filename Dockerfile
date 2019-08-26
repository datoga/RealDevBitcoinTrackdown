FROM golang AS builder

ENV GO111MODULE=on

RUN mkdir -p /go/src/BitcoinTrackdown
WORKDIR /go/src/BitcoinTrackdown

ADD . /go/src/BitcoinTrackdown

WORKDIR /go/src/BitcoinTrackdown

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -s -extldflags "-static"' .

FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir /app 
WORKDIR /app
COPY --from=builder /go/src/BitcoinTrackdown/RealDevBitcoinTrackdown .

CMD ["/app/RealDevBitcoinTrackdown"]