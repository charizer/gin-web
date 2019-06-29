FROM alpine

RUN ls

RUN set -ex \
  && apk add --no-cache ca-certificates

COPY gin-web /opt/gin-web

ENTRYPOINT /opt/gin-web
