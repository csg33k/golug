FROM golang:1.19.2  AS build-stage

LABEL app="build-gb-svc-www"
LABEL REPO="https://github.com/safaci2000/golug"
ENV CGO_ENABLED=0

ENV PROJPATH=/go/src/github.com/safaci2000/golug

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/safaci2000/golug
WORKDIR /go/src/github.com/safaci2000/golug

RUN make linux

# Final Stage
FROM alpine:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="github.com/safaci2000/golug"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

ENV PATH=$PATH:/opt/shiny/bin:/opt/shiny/

WORKDIR /opt/shiny/
## Create appuser
RUN adduser -S shiny -h /opt/shiny/ 

COPY --from=build-stage /go/src/github.com/safaci2000/golug/www_svc_linux /opt/shiny/
RUN \
    apk add dumb-init bash && \
    chmod +x /opt/shiny/www_svc_linux

USER shiny
VOLUME [ "/opt/shiny/conf" ]

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/shiny/www_svc_linux"]