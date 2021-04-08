FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod .

# COPY api/go.sum .

RUN go mod download

RUN go get -u github.com/cosmtrek/air

COPY . .

# EXPOSE 3000

ENTRYPOINT [ "/go/bin/air" ]