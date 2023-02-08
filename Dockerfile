FROM golang:1.19

RUN mkdir -p /app

COPY . /app/

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]