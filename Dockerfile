# 构建前端静态资源
FROM alpine:3.18 AS web
ARG YARN_VERSION=1.22.19-r0
ARG NODEJS_VERSION=18.20.1-r0
RUN apk update && \
    apk add nodejs=${NODEJS_VERSION} yarn=${YARN_VERSION}
ADD web /web
WORKDIR /web
RUN yarn && yarn build

# 构建后端可执行文件
FROM golang:1.22-alpine3.18 AS binary

ARG VERSION=0.6.2

ADD . /go/src/ects
WORKDIR /go/src/ects
COPY --from=web /web/dist /go/src/ects/web/dist
RUN apk update && \
    apk add --no-cache git
RUN go mod tidy && \
    BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ") && \
    GIT_COMMIT=$(git rev-parse HEAD) && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w -X github.com/betterde/ects/internal/build.Version=${VERSION} -X github.com/betterde/ects/internal/build.Build=${BUILD_TIME} -X github.com/betterde/ects/internal/build.Commit=${GIT_COMMIT}" -o "bin/ects_linux" main.go

# 构建运行环境
FROM alpine:3.18

ARG VERSION=0.6.2

LABEL org.opencontainers.image.url="https://betterde.github.io/ects"
LABEL org.opencontainers.image.titile="ECTS"
LABEL org.opencontainers.image.vendor="Betterde Inc."
LABEL org.opencontainers.image.source="https://github.com/betterde/ects"
LABEL org.opencontainers.image.version="${VERSION}"
LABEL org.opencontainers.image.authors="George <george@betterde.com>"
LABEL org.opencontainers.image.created="2024-08-21 20:35:00"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.description="Elastic Crontab System"
LABEL org.opencontainers.image.documentation="https://betterde.github.io/ects"

COPY --from=binary /go/src/ects/bin/ects_linux /usr/local/bin/ects

HEALTHCHECK --interval=60s CMD curl http://localhost/ || exit 1

EXPOSE 9701
CMD ["master"]
ENTRYPOINT ["/usr/local/bin/ects"]
