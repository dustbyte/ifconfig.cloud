FROM golang:1.16.6

WORKDIR /go/src/app
COPY . .

RUN go build .

ENTRYPOINT ["go", "run", "."]
