FROM golang:alpine as build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY api/go.mod .

# COPY go.sum .

RUN go mod download

COPY api/ .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

FROM scratch

COPY --from=build /dist/main /

ENTRYPOINT ["/main"]