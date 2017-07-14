FROM golang:1.8-alpine

WORKDIR /go/src/github.com/opencontrol/compliance-masonry
RUN apk add --no-cache git
COPY . .
RUN go install

ENTRYPOINT ["/go/bin/compliance-masonry"]
