#this is only for development
FROM golang:1.9.2-alpine

ENV ENV_API local
ENV FR_CIRCLE_API_DIR /go/src/github.com/tsrnd/trainning

#install git
RUN apk add --no-cache git mercurial

# install dependency tool
RUN go get -u github.com/golang/dep/cmd/dep

# Fresh for rebuild on code change, no need for production
RUN go get -u github.com/pilu/fresh

# for development, pilu/fresh is used to automatically build the application everytime you save a Go or template file
CMD fresh

EXPOSE 8080 18080