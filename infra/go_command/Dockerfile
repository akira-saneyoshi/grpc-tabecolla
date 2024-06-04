FROM golang:1.22.3-alpine

RUN apk update \
&& apk add alpine-sdk curl \
&& apk --update add tzdata \
&& cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
&& apk del tzdata \
&& rm -rf /var/cache/apk/*

WORKDIR /opt/command_service
COPY ./tabecolla_backend/command_service /opt/command_service
RUN go mod tidy \
  && go install github.com/onsi/ginkgo/v2/ginkgo