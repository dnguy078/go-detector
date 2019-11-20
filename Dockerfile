# alpine does not contain git for dep, should use scratch image and mount the go binaries onto the docker container
FROM golang:1.10.3

# Set go bin which doesn't appear to be set already.
ENV GOBIN /go/bin

ADD . /go/src/github.com/dnguy078/go-detector
WORKDIR /go/src/github.com/dnguy078/go-detector

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure --vendor-only

# TODO make multistage 
