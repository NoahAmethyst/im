FROM golang:1.19.7-alpine3.17

ENV PORT=8080

RUN go env -w GO111MODULE=on \
  && go env -w CGO_ENABLED=0 \
  && go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /opt

COPY ./ .

RUN go mod download

RUN go build
EXPOSE 8080

ENTRYPOINT ["./im"]