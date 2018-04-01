FROM golang:1.10.0-alpine AS build

ARG VCS_REF="unknown"
ARG BUILD_DATE="unknown"

WORKDIR /go/src/github.com/avegao/iot-fronius

RUN apk add --no-cache --update \
        curl \
        git && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY ./ ./

RUN dep ensure && \
    go test ./... -cover && \
    go install \
        -ldflags "-X main.buildDate=$BUILD_DATE -X main.commitHash=$VCS_REF"

########################################################################################################################

FROM alpine:3.7

MAINTAINER "Álvaro de la Vega Olmedilla <alvarodlvo@gmail.com>"

ENV GRPC_VERBOSITY ERROR

RUN addgroup iot-fronius && \
    adduser -D -G iot-fronius iot-fronius

USER iot-fronius

WORKDIR /app

COPY --from=build --chown=iot-fronius:iot-fronius /go/bin/iot-fronius /app/iot-fronius

EXPOSE 50000/tcp

LABEL com.avegao.iot.fronius.vcs_ref=$VCS_REF \
      com.avegao.iot.fronius.build_date=$BUILD_DATE \
      maintainer="Álvaro de la Vega Olmedilla <alvarodlvo@gmail.com>"

ENTRYPOINT ["./iot-fronius"]
