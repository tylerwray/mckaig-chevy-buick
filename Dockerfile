FROM golang:1.10-alpine as base

RUN apk update && \
    apk upgrade && \
    apk add git

WORKDIR /go/src/github.com/tylerwray/red-scare
RUN go get github.com/PuerkitoBio/goquery
COPY . .

FROM base as builder
RUN go build

FROM alpine:3.7
WORKDIR /bin
COPY --from=builder /go/src/github.com/tylerwray/red-scare/red-scare .
ENTRYPOINT [ "red-scare" ]
