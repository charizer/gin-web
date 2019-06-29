FROM golang:1.11.1-alpine
RUN mkdir /build && mkdir -p /var/logs
WORKDIR /build
ADD . /build
RUN sh build.sh
ENTRYPOINT ["./gin-web"]
