FROM golang:1.22.4-alpine
ARG PROTOBUF_VERSION=27.0
EXPOSE 8090

RUN apk update \
&& apk add alpine-sdk \
&& apk --update add tzdata \
&& cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
&& apk del tzdata \
&& rm -rf /var/cache/apk/*

WORKDIR /tmp/protoc
RUN apk add --no-cache unzip curl \
  && curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip -o protoc.zip \
  && unzip protoc.zip \
  && mv /tmp/protoc/bin/* /usr/local/bin/ \
  && mv /tmp/protoc/include /usr/local/include \
  && rm -rf protoc protoc.zip \
  && apk add --no-cache protobuf-dev

WORKDIR /opt/protoc_service
COPY ./tabecolla_backend/protoc_service /opt/protoc_service
RUN go mod tidy \
  && go install google.golang.org/protobuf/cmd/protoc-gen-go \
  && go get -u  google.golang.org/grpc/cmd/protoc-gen-go-grpc \
  && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc