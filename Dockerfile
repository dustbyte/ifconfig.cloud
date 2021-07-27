FROM golang:1.16.6

EXPOSE 80
EXPOSE 8000

WORKDIR /go/src/app
COPY . .

RUN go build .

ENTRYPOINT ["go", "run", "."]
