FROM golang:alpine AS builder
LABEL Omekov Azamat <umekovazamat@gmail.com>
# ENV HTTP_PROXY=http://fgt.halykbank.nb:8080
# ENV HTTPS_PROXY=http://fgt.halykbank.nb:8080
WORKDIR $GOPATH/src/github.com/omekov/online-market/backend
COPY go.mod .
COPY go.sum .
# RUN go get -d -v
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/online-market

FROM scratch
COPY --from=builder /go/bin/online-market /go/bin/online-market
ENTRYPOINT ["/go/bin/online-market"]