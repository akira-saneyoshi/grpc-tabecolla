FROM golang:1.22.3-alpine

RUN apk update \
&& apk add alpine-sdk curl \
&& apk --update add tzdata \
&& cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
&& apk del tzdata \
&& rm -rf /var/cache/apk/*

WORKDIR /opt/query_service
COPY ./tabecolla_backend/query_service /opt/query_service