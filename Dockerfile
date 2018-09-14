FROM golang:1.11.0-alpine3.8
LABEL maintainer="Andreas Peters <support@aventer.biz>"

ENV API_PROXYPORT=10777 
ENV API_PROXYBIND=0.0.0.0
ENV API_URL=http://test/api/v1
ENV API_TOKEN=

COPY  . /src/

RUN apk update && \
    apk add git gcc libc-dev && \
    cd /src/ && \
    go get -d 

RUN cd /src/ && \
    go build -ldflags "-X main.MinVersion=`date -u +%Y%m%d%.H%M%S`" app.go init.go && \
    cp app /  && \
    rm -rf /src

EXPOSE 10777

USER nobody

CMD ["/app"]
