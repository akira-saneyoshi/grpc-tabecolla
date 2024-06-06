FROM golang:1.22.4-alpine

RUN apk update \
&& apk add alpine-sdk curl \
&& apk --update add tzdata \
&& cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
&& apk del tzdata \
&& rm -rf /var/cache/apk/*

WORKDIR /opt/command_service
COPY ./tabecolla_backend/command_service /opt/command_service
RUN go mod tidy \
  && go install github.com/onsi/ginkgo/v2/ginkgo \
  && go install github.com/volatiletech/sqlboiler/v4@latest \
  && go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest