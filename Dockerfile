# Dockerfile used to build the application

ARG arch=amd64
ARG goos=linux

# Build container containing our pre-pulled libraries
FROM golang:alpine AS build

# The golang alpine image is missing git so ensure we have additional tools
RUN apk add --no-cache \
      curl \
      git

# Our build scripts
ADD scripts/ /usr/local/bin/

# Ensure we have the libraries - docker will cache these between builds
RUN get.sh

# Ensure we have the libraries - docker will cache these between builds
#RUN go get -v \
#      github.com/peter-mount/golib/... \
#      gopkg.in/yaml.v2

# ============================================================
# source container contains the source as it exists within the
# repository.
FROM build AS source
WORKDIR /go/src/github.com/peter-mount/calculator
ADD . .

# ============================================================
# Run all tests in a new container so any output won't affect
# the final build.
FROM source as test
ARG skipTest
RUN if [ -z "$skipTest" ] ;then test.sh; fi
