FROM golang:alpine

WORKDIR /go/src/github.com/blyndusk/cofy/api

COPY go.mod .

# COPY api/go.sum .

RUN go mod download

RUN go get -u github.com/cosmtrek/air

COPY . .

ENTRYPOINT [ "/go/bin/air" ]