FROM golang:alpine AS builder
MAINTAINER Omekov Azamat <umekovazamat@gmail.com>


WORKDIR $GOPATH/src/mypackage/market/
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/market

FROM scratch
COPY --from=builder /go/bin/market /go/bin/market
ENTRYPOINT ["/go/bin/market"]