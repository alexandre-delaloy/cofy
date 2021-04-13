FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/github.com/blyndusk/cofy/app

COPY go.mod .

# COPY api/go.sum .

RUN go mod download

RUN go get -u github.com/cosmtrek/air

COPY . .

ENTRYPOINT [ "/go/bin/air" ]