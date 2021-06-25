FROM golang:alpine as builder

# ----- SETUP -----

# Set the image source for ghcr.io
LABEL org.opencontainers.image.source = "https://github.com/blyndusk/cofy"

# Enable Go modules
ENV GO111MODULE=on

# Set the current working with go absolute path
WORKDIR /go/src/github.com/blyndusk/cofy/app

# ----- INSTALL -----

# Copy go.mod + go.sum for install before full copy
COPY app/go.mod .

COPY app/go.sum .

# Download all dependencies
RUN go mod download -x

# ----- COPY + RUN -----

# Copy the source from the current directory to the container
COPY app/ .

# Build app into specific folder
RUN go build -o ./tmp/main ./

# ----- ALPINE -----

FROM alpine

# Copy binary
COPY --from=builder /go/src/github.com/blyndusk/cofy/app/ /cofy/app/

# Set current directory
WORKDIR /cofy/app/

# Use executable
ENTRYPOINT [ "/cofy/app/tmp/main" ]