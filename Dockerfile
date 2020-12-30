FROM golang
LABEL maintainer=idevlab<idevlab@outlook.com>
RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY . $GOPATH/src/funnel/
WORKDIR $GOPATH/src/funnel/

RUN go build .

ENTRYPOINT ["./funnel"]