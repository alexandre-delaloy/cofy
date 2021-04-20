FROM golang:alpine

WORKDIR /go/src/github.com/blyndusk/cofy/app

COPY go.mod .

COPY go.sum .

RUN go mod download

RUN go get -u github.com/cosmtrek/air

COPY . .

ENTRYPOINT [ "/go/bin/air" ]
