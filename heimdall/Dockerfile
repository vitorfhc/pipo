FROM golang:1.16-alpine

WORKDIR /app

COPY . /app

RUN go get -u

RUN go build

CMD ["./heimdall"]