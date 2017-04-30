FROM golang:latest

# Install the  dependencies
RUN go get github.com/gorilla/mux

MAINTAINER jacadenac@unal.edu.co
EXPOSE 8080