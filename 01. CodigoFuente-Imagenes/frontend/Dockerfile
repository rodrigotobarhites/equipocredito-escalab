# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.19.0 as builder

# Create and change to the app directory.
ADD "https://www.random.org/cgi-bin/randbyte?nbytes=10&format=h" skipcache

WORKDIR /app

# Retrieve application dependencies using go modules.
# Allows container builds to reuse downloaded dependencies.
#RUN go clean -i ./

RUN go clean -cache 
COPY go.* ./


RUN go mod download

#RUN go get github.com/ethereum/go-ethereum/ethclient
# Copy local code to the container image.
COPY . ./

# Build the binary.
# -mod=readonly ensures immutable go.mod and go.sum in container builds.

RUN  CGO_ENABLED=0 GOOS=linux go build -mod=mod -v -a -o server

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /server
COPY --from=builder /app/template.html /template.html

# Run the web service on container startup.
CMD ["/server"]

