FROM golang:alpine

WORKDIR /app

COPY services/order/ /app/

RUN go mod tidy
RUN go build -o binary cmd/*.go

ENTRYPOINT ["/app/binary"]