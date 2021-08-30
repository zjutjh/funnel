FROM golang@sha256:504366de834ff066f79a83f394cf3b3221f751d26524e6d8527c8aeec16818a4
LABEL maintainer=idevlab<idevlab@outlook.com> 
LABEL version=2.0.1
RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY . $GOPATH/src/funnel/
WORKDIR $GOPATH/src/funnel/

RUN go build .

ENTRYPOINT ["./funnel"]
