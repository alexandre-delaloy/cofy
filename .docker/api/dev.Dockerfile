FROM golang:alpine

WORKDIR /go/src/github.com/blyndusk/cofy/api

COPY go.mod .

COPY go.sum .

RUN go mod download

RUN go get -u github.com/cosmtrek/air

COPY . .

EXPOSE 3003

ENTRYPOINT [ "/go/bin/air" ]
