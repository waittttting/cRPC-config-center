FROM golang:latest as builder

COPY . /app
WORKDIR /app

RUN go env -w GOPROXY="https://goproxy.cn,direct"
RUN go env -w GO111MODULE=on

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o cx-config-center main.go
FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app /app
RUN echo "Asia/Shanghai" >  /etc/timezone

EXPOSE 8000

CMD ./cx-config-center -config ./config_center/conf/ccs.toml
