# Build the manager binary
FROM golang:1.22 as builder

ARG GITHUB_TOKEN

# TODO: this is temporary to allow access to private repositories.  once this is all public, we
#       no longer need this.
ENV GOPRIVATE=github.com/tbd-paas/*
RUN echo "machine github.com" >> /root/.netrc && \
        echo "login ${GITHUB_TOKEN}" >> /root/.netrc && \
        echo "password x-oauth-basic" >> /root/.netrc && \
        chmod 0600 /root/.netrc

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY internal/ internal/

# Build
RUN CGO_ENABLED=0 go build -a -o platform-bootstrapper main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/platform-bootstrapper .
USER 65532:65532

ENTRYPOINT ["/platform-bootstrapper"]
