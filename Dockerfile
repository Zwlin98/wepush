FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build

FROM ubuntu:20.04
WORKDIR /app
COPY --from=builder /app/wepush .
COPY config.json .
RUN apt-get -y update
RUN apt-get -y install ca-certificates
EXPOSE 9854
CMD ["./wepush"]