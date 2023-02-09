# Stage #1 Build project binary
FROM golang:1.19-alpine AS golang-builder

ARG GOPROXY
ENV GOPROXY ${GOPROXY:-direct}

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PACKAGE github.com/dharanad/journal
ENV BUILD_DIR ${GOPATH}/src/${PACKAGE}

COPY . ${BUILD_DIR}
WORKDIR ${BUILD_DIR}

RUN apk --no-cache add build-base git \
    && make all \
    && cp journal /usr/bin/journal

# Stage #2
FROM alpine

COPY --from=golang-builder /usr/bin/journal /usr/bin/journal

EXPOSE 80
ENTRYPOINT ["/usr/bin/journal"]