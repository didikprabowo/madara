FROM golang

ADD . /go/src/api
RUN go install -i github.com/widnyana/madara
ENTRYPOINT /go/bin/madara

EXPOSE 8080
