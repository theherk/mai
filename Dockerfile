FROM golang:1.13

WORKDIR /go/src/mai
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT ["mai"]
