FROM golang:alpine as builder
COPY . /go/src/echo
WORKDIR /go/src/echo
RUN go build

FROM alpine
MAINTAINER longfangsong@icloud.com
COPY --from=builder /go/src/echo/echo /
WORKDIR /
EXPOSE 80
CMD ["/echo"]