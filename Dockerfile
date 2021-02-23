FROM golang:latest

WORKDIR /go/src/
COPY . .

RUN go get -v
RUN GO111MODULE=auto go build -o app -v

CMD ["/go/src/app"]