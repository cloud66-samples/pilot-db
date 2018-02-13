FROM golang:1.6.4-wheezy

ENV APP_HOME /go/src/app
COPY . $APP_HOME
WORKDIR $APP_HOME

RUN go get -d -v
RUN go build
