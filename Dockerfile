FROM golang:1.13-alpine as builder
# Install some dependencies needed to build go binary
RUN apk add bash ca-certificates git gcc g++ libc-dev

WORKDIR /go/src/github.com/dnguy078/go-detector

COPY go.mod .
COPY go.sum .
RUN go mod download

ADD . /go/src/github.com/dnguy078/go-detector

# RUN go test ./...
# RUN go test ./... -tags=integration

# CGO must be enabled for sqlite :(
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o app main.go
RUN ls

# final image
FROM alpine:3.10.3
RUN apk --no-cache add ca-certificates

WORKDIR /root/

RUN mkdir -p ./schema
RUN mkdir -p ./data

COPY --from=0 /go/src/github.com/dnguy078/go-detector/app .
COPY --from=builder /go/src/github.com/dnguy078/go-detector/schema/  ./schema
COPY --from=builder /go/src/github.com/dnguy078/go-detector/data/  ./data

EXPOSE 3000
RUN ls -lR
CMD ["./app"]