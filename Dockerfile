FROM golang:alpine as builder

WORKDIR /build

COPY . /build/

RUN apk add git && \
    go get -d

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.MinVersion=`date -u +%Y%m%d%.H%M%S` -extldflags \"-static\"" -o main .


FROM alpine
LABEL maintainer="Andreas Peters <support@aventer.biz>"

ENV API_PROXYPORT=10777 
ENV API_PROXYBIND=0.0.0.0
ENV API_URL=http://test/api/v1
ENV API_TOKEN=

RUN apk add --no-cache ca-certificates
RUN adduser -S -D -H -h /app appuser

USER appuser

COPY --from=builder /build/main /app/

EXPOSE 10777

WORKDIR "/app"

CMD ["./main"]
