FROM golang:latest

LABEL maintainer="Abdul <mannanmcc@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8000

RUN go build

CMD ["./helloworld"]
