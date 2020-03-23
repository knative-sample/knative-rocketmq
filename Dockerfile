# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM registry.cn-hangzhou.aliyuncs.com/knative-sample/golang:1.13-alpine3.10 as builder

# Copy local code to the container image.
WORKDIR /go/src/github.com/knative-sample/knative-rocketmq
COPY . .

# Build the command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o knative-rocketmq github.com/knative-sample/knative-rocketmq/cmd


# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM registry.cn-beijing.aliyuncs.com/knative-sample/centos:7.6.1810
COPY --from=builder /go/src/github.com/knative-sample/knative-rocketmq/knative-rocketmq /knative-rocketmq

# Run the web service on container startup.
CMD ["/knative-rocketmq"]
