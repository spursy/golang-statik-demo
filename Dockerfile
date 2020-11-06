# build stage
FROM golang:1.15.3-alpine AS build-env
RUN apk --no-cache add build-base
ENV GOPROXY=https://goproxy.cn
RUN mkdir /tmp/app && \
    apk add curl

COPY . /tmp/app

RUN cd /tmp/app && \
    go mod vendor &&\
#    go generate &&\
    go build -o statik_demo .

FROM alpine:latest
WORKDIR /app
EXPOSE 8080
COPY --from=build-env /tmp/app/statik_demo /app/statik_demo
CMD ["./statik_demo"]
