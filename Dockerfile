FROM quay.io/boundlessgeo/sonar-maven-py3-alpine

RUN apk add --no-cache go protobuf git musl-dev
RUN go get -u github.com/golang/protobuf/protoc-gen-go
ENV GOPATH="/root/go"
ENV GOBIN="${GOPATH}/bin"
ENV PATH="${GOBIN}:${PATH}"

