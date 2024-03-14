FROM golang:1.21-bullseye

RUN go install github.com/beego/bee/v2@latest

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_HOME /go/src/bussiness-auth
RUN sudo mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME/cmd"
EXPOSE 8010
CMD ["bee", "run"]