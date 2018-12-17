# Dockerfile used to build the application

ARG arch=amd64
ARG goos=linux

# Build container containing our pre-pulled libraries
FROM golang:alpine AS build

# The golang alpine image is missing git so ensure we have additional tools
RUN apk add --no-cache \
      curl \
      git \
      tzdata \
      zip

# ============================================================
# source container contains the source as it exists within the
# repository.
FROM build AS source
WORKDIR /work

# Download dependencies before copying any sources then we
# can use the docker cache to limit updates
ADD go.mod .
RUN go mod download

ADD . .

# ============================================================
# Run all tests in a new container so any output won't affect
# the final build.
FROM source as test
ARG skipTest
WORKDIR /work
RUN if [ -z "$skipTest" ] ;\
    then \
      for bin in test;\
      do \
        echo "Testing ${bin}";\
        go test -v github.com/peter-mount/calculator/${bin};\
      done;\
    fi

# ============================================================
FROM source as compiler
WORKDIR /work

# NB: CGO_ENABLED=0 forces a static build
RUN CGO_ENABLED=0 \
    GOOS=${goos} \
    GOARCH=${goarch} \
    GOARM=${goarm} \
    go build \
      -o /dest/usr/local/bin/calculator \
      github.com/peter-mount/calculator/bin

# ============================================================
# Optional stage, upload the binaries as a tar file
FROM compiler AS upload
ARG uploadPath=
ARG uploadCred=
ARG uploadName=
RUN if [ -n "${uploadCred}" -a -n "${uploadPath}" -a -n "${uploadName}" ] ;\
    then \
      cd /dest/usr/local/bin; \
      tar cvzpf /tmp/${uploadName}.tgz * && \
      zip /tmp/${uploadName}.zip * && \
      curl -u ${uploadCred} --upload-file /tmp/${uploadName}.tgz ${uploadPath}/ && \
      curl -u ${uploadCred} --upload-file /tmp/${uploadName}.zip ${uploadPath}/; \
    fi

# ============================================================
# This is the final image
FROM alpine
RUN apk add --no-cache tzdata
COPY --from=compiler /dest/ /
