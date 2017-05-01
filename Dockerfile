
#FROM golang:onbuild
#MAINTAINER jacadenac@unal.edu.co
#EXPOSE 9000 9010


FROM golang:1.8
COPY . /go/src/app
WORKDIR /go/src/app
RUN go-wrapper download
RUN go-wrapper install
RUN go build .
RUN cd phoenix && go build .
RUN cd microservice && go build .
MAINTAINER jacadenac@unal.edu.co
EXPOSE 9000 9010


