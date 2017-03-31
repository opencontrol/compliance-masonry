FROM golang:1.8-alpine
RUN apk add --no-cache git
WORKDIR /go/src/github.com/opencontrol/compliance-masonry
ADD . .
RUN go install
ENTRYPOINT ["/go/bin/compliance-masonry"]
