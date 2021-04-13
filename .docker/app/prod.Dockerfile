FROM golang:alpine as build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY app/go.mod .

# COPY go.sum .

RUN go mod download

COPY app/ .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

FROM scratch

COPY --from=build /dist/main /

ENTRYPOINT ["/main"]