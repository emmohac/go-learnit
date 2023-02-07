FROM golang:latest

RUN mkdir /app

Copy . /app/

WORKDIR /app

Run go build -o main .

CMD ["/app/main"]