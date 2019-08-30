FROM golang:1.12.9-alpine3.10

WORKDIR /go/src/app
COPY *.go ./

RUN apk update && apk add git
RUN go get
RUN go build

CMD ["app"]
