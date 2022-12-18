FROM golang:alpine

WORKDIR /app

COPY account/ /app/

RUN go mod tidy
RUN go build -o binary cmd/*.go

ENTRYPOINT ["/app/binary"]