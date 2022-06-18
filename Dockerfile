FROM golang:1.13.0-alpine3.10
RUN mkdir /sampleapp
ADD . /sampleapp
WORKDIR /sampleapp

